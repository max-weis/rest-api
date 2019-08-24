package app

import (
	"encoding/json"
	"net/http"
)

func handleGet(a *App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, 200, "handleGet Hit!")
	}
}
func handleGetAll(w http.ResponseWriter, r *http.Request, a *App) {
	return func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, 200, "handleGetAll Hit!")
	}}
func handleCreateNew(w http.ResponseWriter, r *http.Request, a *App) {
	return func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, 200, "handleCreateNew Hit!")
	}}
func handleUpdate(w http.ResponseWriter, r *http.Request, a *App) {
	return func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, 200, "handleUpdate Hit!")
	}}
func handleDelete(w http.ResponseWriter, r *http.Request, a *App) {
	return func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, 200, "handleDelete Hit!")
	}
}

// SetRoute bootstraps the routes to the router
func (a *App)SetRoute() {
	a.Router.Methods("GET").Path("/api/books/{id}").Handler(handleGet(a))
	a.Router.Methods("GET").Path("/api/books/{i").Handler(handleGetAll(a))
	a.Router.Methods("GET").Path("/api/books/{id}").Handler(handleGet(a))
	a.Router.Methods("GET").Path("/api/books/{id}").Handler(handleGet(a))
	a.Router.Methods("GET").Path("/api/books/{id}").Handler(handleGet(a))
	a.Router.Methods("GET").Path("/api/books/{id}").Handler(handleGet(a))
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server Error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}
