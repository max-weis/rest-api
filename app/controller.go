package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	// use a dot import to avoid models.Book
	. "gitlab.com/baroprime/prod-rest/models"
)

func handleGet(a *App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		isbn := vars["isbn"]
		book, err := GetBook(a, isbn)
		if err != nil {
			a.Logger.Warn(err)
			respondJSON(w, 400, fmt.Sprintf("Could not find a Book with an ISBN = %s", isbn))
			return
		}
		respondJSON(w, 200, book)
	}
}
func handleGetAll(a *App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := GetAllBooks(a)
		if err != nil {
			a.Logger.Warn(err)
			respondJSON(w, 500, "Could not find any Books")
			return
		}

		respondJSON(w, 200, books)
	}
}
func handleCreateNew(a *App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		book := Book{}
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			respondJSON(w, 404, "Could not read new Book")
			return
		}
		isbn, err := CreateBook(a, book)
		if err != nil {
			respondJSON(w, 404, "Could not create new Book")
			return
		}
		respondJSON(w, 200, fmt.Sprintf("Created new Book with ISBN:%s", isbn))
	}
}
func handleUpdate(a *App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		book := Book{}
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			respondJSON(w, 404, "Could not read the Book")
			return
		}
		_, err = UpdateBook(a, book)
		if err != nil {
			respondJSON(w, 404, "Could not update the Book")
			return
		}
		respondJSON(w, 200, book)
	}
}
func handleDelete(a *App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		isbn := vars["isbn"]

		_, err := DeleteBook(a, isbn)
		if err != nil {
			a.Logger.Warn(err)
			return
		}
		respondJSON(w, 200, fmt.Sprintf("Removed Book with ISBN:%s", isbn))
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
