package app

import (
	"database/sql"
	"gopkg.in/birkirb/loggers.v1"
	"github.com/gorilla/mux"
	"net/http"

 )

type Repo struct {
	DB *sql.DB
 }
 
 func (r Repo) CheckConnection() error {
	return r.DB.Ping()
 }
 
 type App struct {
	Logger loggers.Contextual
	Repo Repo
	Router mux.Router
 }
 
 func (a App) Run () {
	a.Logger.Info("App starting")
	http.ListenAndServe(":8080", &a.Router)
	if err := a.Repo.CheckConnection(); err != nil {
	   panic(err)
	}
	a.Logger.Info("Mission complete")
 }