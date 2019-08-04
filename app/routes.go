package app

import (
	"encoding/json"
	"net/http"
)

//func routes(a App) {
//	a.Router.HandleFunc("/", handlePing).Methods("GET")
//
//	a.Router.HandleFunc("/api/v1/book", func(w http.ResponseWriter, r *http.Request) {
//		json.NewEncoder(w).Encode("handleBookGet Hit!")
//	}).Methods("GET")
//
//	a.Router.HandleFunc("/api/v1/book", func(w http.ResponseWriter, r *http.Request) {
//		json.NewEncoder(w).Encode("handleBookGetList Hit!")
//	}).Methods("GET")
//
//	a.Router.HandleFunc("/api/v1/book", func(w http.ResponseWriter, r *http.Request) {
//		json.NewEncoder(w).Encode("handleBookCreate Hit!")
//	}).Methods("POST")
//
//	a.Router.HandleFunc("/api/v1/book", func(w http.ResponseWriter, r *http.Request) {
//		json.NewEncoder(w).Encode("handleBookUpdate Hit!")
//	}).Methods("PUT")
//
//	a.Router.HandleFunc("/api/v1/book", func(w http.ResponseWriter, r *http.Request) {
//		json.NewEncoder(w).Encode("handleBookDelete Hit!")
//	}).Methods("DELETE")
//}

func handlePing() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("handleBookGet Hit!")
	}
}
