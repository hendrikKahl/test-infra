// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package tm_bot

import (
	"fmt"
	"net/http"

	"github.com/go-logr/logr"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"k8s.io/client-go/rest"

	"github.com/gardener/test-infra/pkg/apis/config"
	"github.com/gardener/test-infra/pkg/testmachinery/controller/watch"
	"github.com/gardener/test-infra/pkg/testmachinery/ghcache"
	"github.com/gardener/test-infra/pkg/tm-bot/github"
	"github.com/gardener/test-infra/pkg/tm-bot/hook"
	"github.com/gardener/test-infra/pkg/tm-bot/tests"
	"github.com/gardener/test-infra/pkg/tm-bot/ui"
	"github.com/gardener/test-infra/pkg/tm-bot/ui/auth"
)

type options struct {
	log        logr.Logger
	restConfig *rest.Config
	cfg        *config.BotConfiguration

	w watch.Watch
}

func NewOptions(log logr.Logger, restConfig *rest.Config, cfg *config.BotConfiguration) *options {
	return &options{
		log:        log,
		restConfig: restConfig,
		cfg:        cfg,
	}
}

func (o *options) setupDashboard(router *mux.Router, runs *tests.Runs) error {
	var (
		authCfg      = o.cfg.Dashboard.Authentication
		authProvider auth.Provider
	)

	switch authCfg.Provider {
	case config.NoAuthProvider:
		authProvider = auth.NewNoAuth()
	case config.DummyAuthProvider:
		authProvider = auth.NewDummyAuth()
	case config.GitHubAuthProvider:
		authProvider = auth.NewGitHubOAuth(o.log.WithName("authentication"), authCfg.GitHub.Hostname,
			authCfg.GitHub.Organization, authCfg.GitHub.OAuth.ClientID, authCfg.GitHub.OAuth.ClientSecret,
			authCfg.GitHub.OAuth.RedirectURL, authCfg.CookieSecret)
	default:
		return fmt.Errorf("no authentication provider with name %s", authCfg.Provider)
	}

	ui.Serve(o.log, runs, o.cfg.Dashboard.UIBasePath, authProvider, router)
	return nil
}

func (o *options) setupGitHubBot(router *mux.Router, runs *tests.Runs) error {
	cfg := o.cfg.GitHubBot
	if !cfg.Enabled {
		return nil
	}
	ghcache.InitGitHubCache(&cfg.GitHubCache)
	ghClient, err := github.NewManager(o.log.WithName("github"), cfg)
	if err != nil {
		return errors.Wrap(err, "unable to initialize github client")
	}
	hooks, err := hook.New(o.log.WithName("hooks"), ghClient, cfg.WebhookSecret, runs)
	if err != nil {
		return errors.Wrap(err, "unable to initialize webhooks handler")
	}

	router.HandleFunc("/events", hooks.HandleWebhook).Methods(http.MethodPost)
	return nil
}

func loggingMiddleware(log logr.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.V(10).Info(r.RequestURI, "method", r.Method)
			next.ServeHTTP(w, r)
		})
	}
}
