package app

import (
	"gitlab.com/baroprime/prod-rest/db"
	"encoding/json"
	"net/http"

	. "gitlab.com/baroprime/prod-rest/models"
)

func handleGet(a *App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, 200, "handleGet Hit!")
	}
}
func handleGetAll(a *App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, 200, "handleGetAll Hit!")
		books,err := GetAllBooks(a)
		if err != nil{
			respondJSON(w,500,"Could not find any Books")
		}

		err = json.NewDecoder()
	}
}
func handleCreateNew(a *App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, 200, "handleCreateNew Hit!")
		book := Book{}
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			respondJSON(w, 404, "Could not read new Book")
		}
		isbn,err = CreateBook(a)
		if err != nil{
			respondJSON(w, 404, "Could not create new Book")
		}
		respondJSON(w,200,fmt.Sprintf("Created new Book with ISBN:%s",isbn))
	}
}
func handleUpdate(a *App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, 200, "handleUpdate Hit!")
	}
}
func handleDelete(a *App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, 200, "handleDelete Hit!")
	}
}

// SetRoute bootstraps the routes to the router
func (a *App) SetRoute() {
	a.Router.Methods("GET").Path("/api/books/{isbn}").Handler(handleGet(a))
	a.Router.Methods("GET").Path("/api/books").Handler(handleGetAll(a))
	a.Router.Methods("POST").Path("/api/books").Handler(handleCreateNew(a))
	a.Router.Methods("PUT").Path("/api/books").Handler(handleUpdate(a))
	a.Router.Methods("DELETE").Path("/api/books/{isbn}").Handler(handleDelete(a))
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
