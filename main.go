package main

import (
	"github.com/go-nop/ginstarter/cmd"
	"github.com/go-nop/ginstarter/config"
	"github.com/go-nop/ginstarter/pkg/log"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	// init logger
	log.NewLogger(log.ZapDriver)

	// Muat file .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// init config
	config := config.New(
		config.WithAppConfig(),
		config.WithDBConfig(),
	)

	// root command
	rootCmd := &cobra.Command{
		Use:   "http:start",
		Short: "Start HTTP server",
		Run: func(command *cobra.Command, args []string) {
			err := cmd.NewApp(config)
			if err != nil {
				log.Error("failed to start HTTP server", map[string]interface{}{"error": err})
			}
		},
	}

	// Migrate command
	migrateCmd := cmd.MigrateCommand(config)
	rootCmd.AddCommand(migrateCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal("failed to execute command", map[string]interface{}{"error": err})
	}
}
