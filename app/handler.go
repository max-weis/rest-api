package app

import (
	"net/http"
	"encoding/json"
)

func (a *App) handleBookGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		json.NewEncoder(w).Encode("handleBookGet Hit!")
	}
}

func (a *App) handleBookGetList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		json.NewEncoder(w).Encode("handleBookGetList Hit!")
	}
}

func (a *App) handleBookCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		json.NewEncoder(w).Encode("handleBookCreate Hit!")
	}
}

func (a *App) handleBookUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		json.NewEncoder(w).Encode("handleBookUpdate Hit!")
	}
}

func (a *App) handleBookDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		json.NewEncoder(w).Encode("handleBookDelete Hit!")
	}
}
