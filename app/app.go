package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	// postgres dialect
	_ "github.com/lib/pq"

	mapper "github.com/birkirb/loggers-mapper-logrus"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gopkg.in/birkirb/loggers.v1"
)

// App is the main part of the app
type App struct {
	Config Config
	Logger loggers.Contextual
	DB     *sql.DB
	Router *mux.Router
}

// NewLogger creates a new logger
func newLogger() loggers.Contextual {
	l := logrus.New()
	l.Out = os.Stdout
	l.Level = logrus.InfoLevel
	l.SetFormatter(&logrus.JSONFormatter{})
	return mapper.NewLogger(l)
}

// NewDB creates a new DB connection
func newDB(c Config) *sql.DB {
	db, err := sql.Open("postgres", c.getConnectionString())
	if err != nil {
		panic(err)
	}
	return db
}

// NewRouter creates a new router
func newRouter() *mux.Router {
	return mux.NewRouter()
}

// NewApp creates an app object
func NewApp() App {
	return App{}
}

// Init initalizes the app
func (a *App) Init(c Config) {
	a.Config = c
	a.DB = newDB(a.Config)
	a.Logger = newLogger()
	a.Router = newRouter()
}

// Run starts the app
func (a *App) Run(port string) {
	a.Logger.Info("App starting")
	a.SetRoute()

	http.ListenAndServe(port, a.Router)
	a.Logger.Info("Mission complete")
}

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
		panic(fmt.Sprintf("Could not set %s", key))
	}
	return val
}

func (c Config) getConnectionString() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		c.dbUser, c.dbPass, c.dbHost, c.dbPort, c.dbName,
	)
}
