package helpers

import (
	"os"
)

// Getenv gets an environment variable or returns a fallback value.
func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
