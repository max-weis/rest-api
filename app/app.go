package app

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/birkirb/loggers.v1"
)

// App is the main part of the app
type App struct {
	Logger loggers.Contextual
	DB     *sql.DB
	Router mux.Router
}

// NewLogger creates a new logger
func (a *App) NewLogger() {

}

// NewDB creates a new DB connection
func (a *App) NewDB() {

}

// NewRouter creates a new router
func (a *App) NewRouter() {

}

// Run starts the app
func (a *App) Run(port string) {
	a.Logger.Info("App starting")

	http.ListenAndServe(port, &a.Router)
	a.Logger.Info("Mission complete")
}
