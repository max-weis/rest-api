package app

import (
	"database/sql"
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

// SetRoute bootstraps the routes to the router
func (a *App) SetRoute() {
	a.Router.Methods("GET").Path("/api/books/{isbn}").Handler(handleGet(a))
	a.Router.Methods("GET").Path("/api/books").Handler(handleGetAll(a))
	a.Router.Methods("POST").Path("/api/books").Handler(handleCreateNew(a))
	a.Router.Methods("PUT").Path("/api/books").Handler(handleUpdate(a))
	a.Router.Methods("DELETE").Path("/api/books/{isbn}").Handler(handleDelete(a))
	a.Logger.Info("Routes initialized")
}

// NewApp creates an app object
func NewApp() App {
	return App{}
}

// Init initializes the app
func (a *App) Init(c Config) {
	a.Logger = newLogger()
	a.Config = c
	a.DB = newDB(a.Config)
	a.Logger.Info("Config loaded")
	a.Router = newRouter()
}

// Run starts the app
func (a *App) Run(port string) {
	a.Logger.Info("App starting")
	a.SetRoute()

	http.ListenAndServe(port, a.Router)
	a.Logger.Info("Mission complete")
}
