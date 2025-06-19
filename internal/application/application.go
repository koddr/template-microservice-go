package application

import (
	"log/slog"
	"os"

	"github.com/koddr/template-microservice-go/internal/attachments"
	"github.com/koddr/template-microservice-go/internal/config"
	"github.com/koddr/template-microservice-go/internal/database"
)

// Application contains DB connection and other dependencies for application.
type Application struct {
	Attachments *attachments.Attachments
	Config      *config.Config
	Database    *database.Database
	Logger      *slog.Logger
}

// New returns a new instance of Application.
func New(att *attachments.Attachments, cfg *config.Config, db *database.Database) *Application {
	return &Application{
		Attachments: att,
		Config:      cfg,
		Database:    db,
		Logger:      slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}
