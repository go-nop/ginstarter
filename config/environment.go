package config

const (
	// EnvAppName is the environment variable for the application name.
	EnvAppName = "APP_NAME"
	// EnvAppEnviron is the environment variable for the application environment.
	EnvAppEnviron = "APP_ENV"
	// EnvAppTimezone is the environment variable for the application timezone.
	EnvAppTimezone = "APP_TIMEZONE"
	// EnvAppDebugMode is the environment variable for the application debug mode.
	EnvAppDebugMode = "APP_DEBUG"

	// EnvDBDriver is the environment variable for the database driver.
	EnvDBDriver = "DB_DRIVER"
	// EnvDBHost is the environment variable for the database host.
	EnvDBHost = "DB_HOST"
	// EnvDBPort is the environment variable for the database port.
	EnvDBPort = "DB_PORT"
	// EnvDBUser is the environment variable for the database user.
	EnvDBUser = "DB_USER"
	// EnvDBPassword is the environment variable for the database password.
	EnvDBPassword = "DB_PASSWORD"
	// EnvDBName is the environment variable for the database name.
	EnvDBName = "DB_NAME"
	// EnvDBSSLMode is the environment variable for the database SSL mode.
	EnvDBSSLMode = "DB_SSL_MODE"
	// EnvDBMaxIdleConns is the environment variable for the maximum number of idle connections in the connection pool.
	EnvDBMaxIdleConns = "DB_MAX_IDLE_CONNS"
	// EnvDBMaxOpenConns is the environment variable for the maximum number of open connections to the database.
	EnvDBMaxOpenConns = "DB_MAX_OPEN_CONNS"
	// EnvDBMaxLifetime is the environment variable for the maximum amount of time a connection may be reused.
	EnvDBMaxLifetime = "DB_MAX_LIFETIME"
	// EnvDBMaxIdleTime is the environment variable for the maximum amount of time a connection may be idle before being closed.
	EnvDBMaxIdleTime = "DB_MAX_IDLE_TIME"
)
