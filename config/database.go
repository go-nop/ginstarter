package config

import (
	"time"

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

	SSLMode      string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  time.Duration
	MaxIdleTime  time.Duration
	TimeZone     string
}

// WithDBConfig sets the database configuration.
func WithDBConfig() Option {
	toTimeDuration := func(value int) time.Duration {
		return time.Duration(value) * time.Second
	}

	return func(c *Config) {
		c.DB = DB{
			Driver:       configurator.GetEnv(EnvDBDriver, consts.DefaultDBDriver),
			Host:         configurator.GetEnv(EnvDBHost, consts.DefaultDBHost),
			Port:         configurator.GetEnv(EnvDBPort, consts.DefaultDBPort),
			User:         configurator.GetEnv(EnvDBUser, consts.DefaultDBUser),
			Password:     configurator.GetEnv(EnvDBPassword, consts.DefaultDBPassword),
			Name:         configurator.GetEnv(EnvDBName, consts.DefaultDBName),
			SSLMode:      configurator.GetEnv(EnvDBSSLMode, consts.DefaultDBSSLMode),
			MaxIdleConns: configurator.GetEnvAsInt(EnvDBMaxIdleConns, consts.DefaultDBMaxIdleConns),
			MaxOpenConns: configurator.GetEnvAsInt(EnvDBMaxOpenConns, consts.DefaultDBMaxOpenConns),
			MaxLifetime:  toTimeDuration(configurator.GetEnvAsInt(EnvDBMaxLifetime, consts.DefaultDBMaxLifetime)),
			MaxIdleTime:  toTimeDuration(configurator.GetEnvAsInt(EnvDBMaxIdleTime, consts.DefaultDBMaxIdleTime)),
			TimeZone:     configurator.GetEnv("DB_TIMEZONE", consts.DefaultAppTimezone),
		}
	}
}
