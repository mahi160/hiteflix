package routers

import (
	"github.com/gorilla/mux"
	"github.com/mahi160/hiteflix/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.MarkWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controllers.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/delete-all-movies", controllers.DeleteAllMovies).Methods("DELETE")
	return router
}
