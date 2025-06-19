package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/koddr/template-microservice-go/internal/config"
)

// Database contains DB connection and other dependencies for application.
type Database struct {
	Connection *pgxpool.Pool
}

// New returns a new instance of DB connection.
func New(c *config.Config) (*Database, error) {
	// Create pool connection to DB.
	connection, err := pgxpool.New(context.Background(), c.Storage.URL)
	if err != nil {
		return nil, err
	}

	return &Database{
		Connection: connection,
	}, nil
}
