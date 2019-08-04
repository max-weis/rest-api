package app

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/birkirb/loggers.v1"
)

type Repo struct {
	DB *sql.DB
}

func (r Repo) CheckConnection() error {
	return r.DB.Ping()
}

type App struct {
	Logger loggers.Contextual
	Repo   Repo
	Router mux.Router
}

func (a App) Run(port string) {
	a.Logger.Info("App starting")
	routes(&a.Router)
	if err := a.Repo.CheckConnection(); err != nil {
		panic(err)
	}
	http.ListenAndServe(port, &a.Router)
	a.Logger.Info("Mission complete")
}
