package attachments

import (
	"embed"
)

var (
	//go:embed migrations/*.sql
	migrationFiles embed.FS

	//go:embed queries/*.sql
	queryFiles embed.FS
)

// Attachments represents struct for embed files.
type Attachments struct {
	Migrations, Queries embed.FS
}

// New creates a new collection with embed files by Attachments struct.
func New() *Attachments {
	return &Attachments{
		Migrations: migrationFiles,
		Queries:    queryFiles,
	}
}
