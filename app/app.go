package app

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/birkirb/loggers.v1"
)

// Repo contains the db
type Repo struct {
	DB *sql.DB
}

// CheckConnection checks if there is a connection to the db
func (r Repo) CheckConnection() error {
	return r.DB.Ping()
}

// App is the main part of the app
type App struct {
	Logger loggers.Contextual
	Repo   Repo
	Router mux.Router
}

// Run starts the app
func (a App) Run(port string) {
	a.Logger.Info("App starting")

	a.Router.HandleFunc("/", handlePing()).Methods("GET")

	if err := a.Repo.CheckConnection(); err != nil {
		panic(err)
	}
	http.ListenAndServe(port, &a.Router)
	a.Logger.Info("Mission complete")
}
