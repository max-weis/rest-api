package app

import (
	"encoding/json"
	"net/http"
)

func handlePing(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("PONG")
}

func handleBookGet(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("handleBookGet Hit!")
}

func handleBookGetList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("handleBookGetList Hit!")
}

func handleBookCreate(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("handleBookCreate Hit!")
}

func handleBookUpdate(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("handleBookUpdate Hit!")
}

func handleBookDelete(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("handleBookDelete Hit!")
}
