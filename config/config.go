package config

// AppConfig is the configuration for the application.
type Config struct {
	App App
	DB  DB
}

// Option is a function that configures the Configurator.
type Option func(*Config)

// New creates a new AppConfig with the provided options.
func New(options ...Option) *Config {
	config := &Config{}
	for _, option := range options {
		option(config)
	}
	return config
}
