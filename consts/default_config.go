package consts

// Default configuration values for the application.
const (
	// DefaultAppName is the default name of the application.
	DefaultAppName = "go-nop-app"
	// DefaultAppEnviron is the default environment the application is running in. (e.g. local, development, production)
	DefaultAppEnviron = "local"
	// DefaultAppTimezone is the default timezone of the application.
	DefaultAppTimezone = "UTC"
	// DefaultAppPort is the default port the application listens on.
	DefaultAppPort = "8080"
	// DefaultAppDebugMode is the default debug mode of the application.
	DefaultAppDebugMode = false
)

// Default configuration values for the database.
const (
	// DefaultDBDriver is the default database driver.
	DefaultDBDriver = "postgres"
	// DefaultDBHost is the default database host.
	DefaultDBHost = "localhost"
	// DefaultDBPort is the default database port.
	DefaultDBPort = "5432"
	// DefaultDBUser is the default database user.
	DefaultDBUser = "postgres"
	// DefaultDBPassword is the default database password.
	DefaultDBPassword = "password"
	// DefaultDBName is the default database name.
	DefaultDBName = "ginstarter"
	// DefaultDBSSLMode is the default database SSL mode.
	DefaultDBSSLMode = "disable"
	// DefaultDBMaxIdleConns is the default maximum number of idle connections in the connection pool.
	DefaultDBMaxIdleConns = 10
	// DefaultDBMaxOpenConns is the default maximum number of open connections to the database.
	DefaultDBMaxOpenConns = 100
	// DefaultDBMaxLifetime is the default maximum amount of time a connection may be reused.
	DefaultDBMaxLifetime = 0
	// DefaultDBMaxIdleTime is the default maximum amount of time a connection may be idle before being closed.
	DefaultDBMaxIdleTime = 0
)
