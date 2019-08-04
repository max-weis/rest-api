package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func routes(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("PONG")
	}).Methods("GET")
	r.HandleFunc("/api/v1/book", handleBookGet).Methods("GET")
	r.HandleFunc("/api/v1/book", handleBookGetList).Methods("GET")
	r.HandleFunc("/api/v1/book", handleBookCreate).Methods("POST")
	r.HandleFunc("/api/v1/book", handleBookUpdate).Methods("PUT")
	r.HandleFunc("/api/v1/book", handleBookDelete).Methods("DELETE")
}
