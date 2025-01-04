package config

import (
	"github.com/go-nop/ginstarter/consts"
	configurator "github.com/go-nop/ginstarter/pkg/config"
)

// DB is the configuration for the database.
type DB struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// WithDBConfig sets the database configuration.
func WithDBConfig() Option {
	return func(c *Config) {
		c.DB = DB{
			Driver:   configurator.GetEnv("DB_DRIVER", consts.DefaultDBDriver),
			Host:     configurator.GetEnv("DB_HOST", consts.DefaultDBHost),
			Port:     configurator.GetEnv("DB_PORT", consts.DefaultDBPort),
			User:     configurator.GetEnv("DB_USER", consts.DefaultDBUser),
			Password: configurator.GetEnv("DB_PASSWORD", consts.DefaultDBPassword),
			Name:     configurator.GetEnv("DB_NAME", consts.DefaultDBName),
		}
	}
}
