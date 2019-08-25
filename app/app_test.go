package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	logger := newLogger()
	assert.NotNil(t, logger)
}

func TestNewRouter(t *testing.T) {
	router := newRouter()
	assert.NotNil(t, router)
}

func TestNewApp(t *testing.T) {
	app := NewApp()

	assert.Equal(t, app.Config.dbUser, "")
	assert.Equal(t, app.Config.dbPass, "")
	assert.Equal(t, app.Config.dbHost, "")
	assert.Equal(t, app.Config.dbPort, "")
	assert.Equal(t, app.Config.dbName, "")

	assert.Nil(t, app.Logger)
	assert.Nil(t, app.DB)
	assert.Nil(t, app.Router)
}

func TestNewConfig(t *testing.T) {
	config := NewConfig()

	assert.Equal(t, config.dbUser, "")
	assert.Equal(t, config.dbPass, "")
	assert.Equal(t, config.dbHost, "")
	assert.Equal(t, config.dbPort, "")
	assert.Equal(t, config.dbName, "")
}

func TestGetConnectionString(t *testing.T) {
	config := Config{
		dbUser: "postgres",
		dbPass: "postgres",
		dbHost: "db",
		dbPort: "5432",
		dbName: "app",
	}

	assert.Equal(
		t,
		config.getConnectionString(),
		"user=postgres password=postgres host=db port=5432 dbname=app sslmode=disable",
	)
}
