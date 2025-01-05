package db

import (
	"fmt"

	"github.com/go-nop/ginstarter/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	// mySQLDriver is the MySQL driver.
	mySQLDriver = "mysql"
	// postgreSQLDriver is the PostgreSQL driver.
	postgreSQLDriver = "postgres"
)

// dsnOptions is a struct that contains the options for the DSN.
type dsnOptions struct {
	driver   string
	host     string
	port     string
	user     string
	password string
	dbName   string
	sslMode  string
	timeZone string
}

// newDSN creates a new DSN.
func newDSN(cfg *config.DB) dsnOptions {
	return dsnOptions{
		driver:   cfg.Driver,
		host:     cfg.Host,
		port:     cfg.Port,
		user:     cfg.User,
		password: cfg.Password,
		dbName:   cfg.Name,
		sslMode:  cfg.SSLMode,
		timeZone: cfg.TimeZone,
	}
}

// toGormDialector converts the DSN to a GORM dialector.
func (d dsnOptions) toGormDialector() gorm.Dialector {
	switch d.driver {
	case mySQLDriver:
		return mysql.Open(d.toMySQLDSN())
	case postgreSQLDriver:
		return postgres.Open(d.toPostgreSQLDSN())
	default:
		return nil
	}
}

// toMySQLDSN converts the DSN to a MySQL DSN.
func (d dsnOptions) toMySQLDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.user, d.password, d.host, d.port, d.dbName)
}

// toPostgreSQLDSN converts the DSN to a PostgreSQL DSN.
func (d dsnOptions) toPostgreSQLDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s", d.host, d.port, d.user, d.password, d.dbName, d.sslMode, d.timeZone)
}
