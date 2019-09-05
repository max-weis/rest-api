package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/baroprime/prod-rest/db"
)

func handleGet(a *App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		isbn := vars["isbn"]
		book, err := db.GetBook(a.DB, isbn)
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
		books, err := db.GetAllBooks(a.DB)
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
		isbn, err := db.CreateBook(a.DB, book)
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
		_, err = db.UpdateBook(a.DB, book)
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

		_, err := db.DeleteBook(a.DB, isbn)
		if err != nil {
			a.Logger.Warn(err)
			return
		}
		respondJSON(w, 200, fmt.Sprintf("Removed Book with ISBN:%s", isbn))
	}
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
