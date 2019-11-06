package app

import (
	"fmt"
	"os"
)

// Config stores all env vars
type Config struct {
	dbUser string
	dbPass string
	dbHost string
	dbPort string
	dbName string
}

// NewConfig inits a Config file
func NewConfig() Config {
	return Config{
		dbUser: getEnvVar("DB_USER"),
		dbPass: getEnvVar("DB_PASS"),
		dbHost: getEnvVar("DB_HOST"),
		dbPort: getEnvVar("DB_PORT"),
		dbName: getEnvVar("DB_NAME"),
	}
}

func getEnvVar(key string) string {
	val := os.Getenv(key)
	if val == "" {
		fmt.Fprintf(os.Stderr, "Could not find env %s\n", key)
	}
	return val
}

func (c Config) getConnectionString() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		c.dbUser, c.dbPass, c.dbHost, c.dbPort, c.dbName,
	)
}
