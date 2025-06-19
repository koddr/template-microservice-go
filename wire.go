//go:build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/koddr/template-microservice-go/internal/application"
	"github.com/koddr/template-microservice-go/internal/attachments"
	"github.com/koddr/template-microservice-go/internal/config"
	"github.com/koddr/template-microservice-go/internal/database"
)

// initializeApplication provides the dependency injection process by the "google/wire" package.
func initializeApplication() (*application.Application, error) {
	wire.Build(attachments.New, config.New, database.New, application.New)
	return &application.Application{}, nil
}
