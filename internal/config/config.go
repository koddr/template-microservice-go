package config

import (
	"os"

	"github.com/koddr/template-microservice-go/internal/helpers"
)

// Config represents struct for an app configuration.
type Config struct {
	Server  *Server
	Storage *Storage
}

// Server represents struct for API server.
type Server struct {
	BasicAuthUsername, BasicAuthPassword, Port string
}

// Storage represents struct for DB.
type Storage struct {
	URL string
}

// New creates a new config.
func New() *Config {
	return &Config{
		Server: &Server{
			BasicAuthUsername: os.Getenv("API_SERVER_AUTH_USERNAME"),
			BasicAuthPassword: os.Getenv("API_SERVER_AUTH_PASSWORD"),
			Port:              helpers.Getenv("API_SERVER_PORT", "8080"),
		},
		Storage: &Storage{
			URL: os.Getenv("API_SERVER_DATABASE_URL"),
		},
	}
}
