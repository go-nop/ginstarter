package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-nop/ginstarter/config"
	"github.com/go-nop/ginstarter/consts"
	"github.com/go-nop/ginstarter/internal/db"
	"github.com/go-nop/ginstarter/internal/repositories"
	"github.com/go-nop/ginstarter/internal/route"
	"github.com/go-nop/ginstarter/internal/services"
	configurator "github.com/go-nop/ginstarter/pkg/config"
	"github.com/go-nop/ginstarter/pkg/log"
)

// NewApp creates a new application.
func NewApp(config *config.Config) error {
	// init app configuration

	time.Local = time.FixedZone(config.App.Timezone, 0)

	// init database connection
	dbConn, err := db.NewDBConnection(&config.DB)
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	// init repository
	repo := repositories.NewRepositories(dbConn)

	// init service
	svc := services.NewServices(repo)

	// init router
	r := route.NewRouter(
		route.WithAppConfig(&config.App),
		route.WithRepositories(repo),
		route.WithServices(svc),
	).Init()

	log.Info("HTTP server is starting...")

	appPort := configurator.GetEnv("APP_PORT", consts.DefaultAppPort)

	server := http.Server{
		Addr:         ":" + appPort,
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	server.SetKeepAlivesEnabled(false)

	serverErr := make(chan error, 1)

	go func() {
		log.Infof("HTTP server is running on port: %s", appPort)
		serverErr <- server.ListenAndServe()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-serverErr:
		if err != nil {
			log.Error("failed to start HTTP server", map[string]interface{}{"error": err})
			return err
		}
	case <-stop:
		log.Info("HTTP server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Error("failed to shutdown HTTP server", map[string]interface{}{"error": err})
			return err
		}

		log.Info("HTTP server is stopped")
	}

	return nil
}
