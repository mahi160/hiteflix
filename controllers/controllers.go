package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mahi160/hiteflix/db"
	"github.com/mahi160/hiteflix/models"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := db.GetAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	db.InsertOne(movie)
	json.NewEncoder(w).Encode(movie)
}
func MarkWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	db.UpdateOne(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
