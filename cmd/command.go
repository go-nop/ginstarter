package cmd

import (
	"github.com/go-nop/ginstarter/config"
	"github.com/go-nop/ginstarter/internal/db"
	"github.com/go-nop/ginstarter/pkg/log"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

const dirMigrations = "internal/db/migrations"

func MigrateCommand(config *config.Config) *cobra.Command {
	migrateCmd := &cobra.Command{
		Use:   "db:migrate",
		Short: "Run database migrations",
		Run: func(command *cobra.Command, args []string) {
			log.Info("Running database migrations")

			// init database connection
			dbConn, err := db.NewDBConnection(&config.DB)
			if err != nil {
				log.Fatal("failed to connect to database")
			}

			sqlDB, err := dbConn.DB()
			if err != nil {
				log.Fatal("failed to get database connection")
			}

			// Run migrations
			err = goose.Up(sqlDB, dirMigrations, goose.WithAllowMissing())
			if err != nil {
				log.Error(err.Error())
			}
		},
	}

	return migrateCmd
}
