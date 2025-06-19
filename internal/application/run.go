package application

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Run runs the application.
func (a *Application) Run() error {
	// Define migrations folder from embedded filesystem.
	migrationsFolder, err := iofs.New(a.Attachments.Migrations, "migrations")
	if err != nil {
		a.Logger.Error("failed to read migrations from the embedded filesystem", "details", err.Error())
		return err
	}

	// Create migrations instance.
	migrations, err := migrate.NewWithSourceInstance(
		"iofs", migrationsFolder,
		fmt.Sprintf("%s?sslmode=disable&search_path=public", a.Config.Storage.URL),
	)
	if err != nil {
		a.Logger.Error("failed database connection to setup migrations", "details", err.Error())
		return err
	}

	// Run migrations.
	if err := migrations.Up(); err != nil && err != migrate.ErrNoChange {
		a.Logger.Error("failed to apply migrations", "details", err.Error())
		return err
	} else if err == migrate.ErrNoChange {
		a.Logger.Info("no new migrations to apply")
	} else {
		a.Logger.Info("all new migrations successfully applied")
	}

	// Create a new server instance with options from environment variables.
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", a.Config.Server.Port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
		Handler:      a.router(), // use the HttpRouter instance with session manager
	}

	// Log the start of the application.
	a.Logger.Info("starting API server", "port", a.Config.Server.Port)

	// Create a context with cancel to manage server lifecycle.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the server in a separate goroutine.
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.Logger.Error("failed to start API server", "error", err)
			cancel() // Cancel the context to trigger shutdown.
		}
	}()

	// Create a channel to listen for OS signals.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Wait for an OS signal or context cancellation.
	select {
	case <-stop:
		a.Logger.Info("received shutdown signal")
	case <-ctx.Done():
		a.Logger.Info("context canceled, shutting down server")
	}

	// Create a context with a timeout for graceful shutdown.
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	// Shutdown the server gracefully.
	if err := server.Shutdown(shutdownCtx); err != nil {
		a.Logger.Error("failed to gracefully shutdown API server", "error", err)
		return err
	}

	a.Logger.Info("gracefully shutdown API server")
	return nil
}
