package router

import (
	"movie_app/controller"
	"github.com/gorilla/mux"
)


func Init() *mux.Router{
	r := mux.NewRouter()

	r.HandleFunc("/movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", controller.GetMovieByID).Methods("GET")
	r.HandleFunc("/movie", controller.AddMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", controller.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	r.HandleFunc("/movies", controller.DeleteAllMovies).Methods("DELETE")
	
	return r
}