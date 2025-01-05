package db

import (
	"github.com/go-nop/ginstarter/config"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// NewDBConnection creates a new DB connection.
func NewDBConnection(cfg *config.DB) (*gorm.DB, error) {
	dsn := newDSN(cfg)

	dialector := dsn.toGormDialector()

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to open database connection")
	}

	db, err = setConnectionPool(db, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set connection pool")
	}

	// TODO: implement database replication

	return db, nil
}

// setConnectionPool sets the connection pool settings.
func setConnectionPool(db *gorm.DB, cfg *config.DB) (*gorm.DB, error) {
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.MaxLifetime)
	sqlDB.SetConnMaxIdleTime(cfg.MaxIdleTime)

	return db, nil
}
