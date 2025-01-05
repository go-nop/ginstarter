package route

import (
	"github.com/go-nop/ginstarter/config"
	"github.com/go-nop/ginstarter/internal/repositories"
	"github.com/go-nop/ginstarter/internal/services"
)

// Option is a function that configures the Router.
type Option func(*Router)

// WithAppConfig sets the app configuration.
func WithAppConfig(appConfig *config.App) Option {
	return func(r *Router) {
		r.appConfig = appConfig
	}
}

// WithRepositories sets the repositories.
func WithRepositories(repositories *repositories.Repositories) Option {
	return func(r *Router) {
		r.repositories = repositories
	}
}

// WithServices sets the services.
func WithServices(services *services.Services) Option {
	return func(r *Router) {
		r.services = services
	}
}
