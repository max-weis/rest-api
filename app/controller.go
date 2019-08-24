package app

import (
	"encoding/json"
	"net/http"
)

func handleGet(a *App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, 200, []byte("handleGet Hit!"))
	}
}
func handleGetAll(w http.ResponseWriter, r *http.Request, a *App) {
	respondJSON(w, 200, []byte("handleGetAll Hit!"))
}
func handleCreateNew(w http.ResponseWriter, r *http.Request, a *App) {
	respondJSON(w, 200, []byte("handleGetCreate Hit!"))
}
func handleUpdate(w http.ResponseWriter, r *http.Request, a *App) {
	respondJSON(w, 200, []byte("handleGetUpdate Hit!"))
}
func handleDelete(w http.ResponseWriter, r *http.Request, a *App) {
	respondJSON(w, 200, []byte("handleGetDelete Hit!"))
}

// SetRoutes bootstraps the routes to the router
func (a *App) SetRoutes() {
	a.Router.Methods("GET").Path("/api/books").Handler(handleGet(a))
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
