package config

import (
	"github.com/go-nop/ginstarter/consts"
	configurator "github.com/go-nop/ginstarter/pkg/config"
)

// App is the configuration for the application.
type App struct {
	Name     string
	Environ  string
	Timezone string
}

// WithAppConfig sets the application configuration.
func WithAppConfig() Option {
	return func(c *Config) {
		c.App = App{
			Name:     configurator.GetEnv("APP_NAME", consts.DefaultAppName),
			Environ:  configurator.GetEnv("APP_ENVIRON", consts.DefaultAppEnviron),
			Timezone: configurator.GetEnv("APP_TIMEZONE", consts.DefaultAppTimezone),
		}
	}
}
