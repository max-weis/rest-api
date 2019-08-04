package db

import (
	"fmt"
	"os"
)

type Config struct {
	dbUser string
	dbPass string
	dbHost string
	dbPort string
	dbName string
}

func GetConfig() Config {
	return Config{
		dbUser: getEnvVar("DB_USER"),
		dbPass: getEnvVar("DB_PASS"),
		dbHost: getEnvVar("DB_HOST"),
		dbPort: getEnvVar("DB_PORT"),
		dbName: getEnvVar("DB_NAME"),
	}
}
func (c Config) GetConnectionString() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		c.dbUser,
		c.dbPass,
		c.dbHost,
		c.dbPort,
		c.dbName,
		"disable",
	)
}
func getEnvVar(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(
			fmt.Sprintf("env var not set %s", key),
		)
	}
	return val
}
