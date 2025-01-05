package config

import (
	"github.com/go-nop/ginstarter/consts"
	configurator "github.com/go-nop/ginstarter/pkg/config"
)

// App is the configuration for the application.
type App struct {
	Name      string
	Environ   string
	Timezone  string
	DebugMode bool
}

// WithAppConfig sets the application configuration.
func WithAppConfig() Option {
	return func(c *Config) {
		c.App = App{
			Name:      configurator.GetEnv(EnvAppName, consts.DefaultAppName),
			Environ:   configurator.GetEnv(EnvAppEnviron, consts.DefaultAppEnviron),
			Timezone:  configurator.GetEnv(EnvAppTimezone, consts.DefaultAppTimezone),
			DebugMode: configurator.GetEnvAsBool(EnvAppDebugMode, consts.DefaultAppDebugMode),
		}
	}
}
