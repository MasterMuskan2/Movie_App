package controller

import (
	"math/rand"
	"encoding/json"
	"movie_app/database"
	"movie_app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Get all the movies present in the db

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(database.Movies)
}

// Get a movie by it's ID

func GetMovieByID(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)

	for _, value := range database.Movies{
		if value.Id == params["id"]{
			json.NewEncoder(w).Encode(value)
			return
		}
	}
	json.NewEncoder(w).Encode("Movie not found with the given ID!!")
}

// Add one movie to the db

func AddMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie model.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(100000))
	database.Movies = append(database.Movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// Update a movies from the db

func UpdateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, value := range database.Movies{
		if params["id"] == value.Id{
			database.Movies = append(database.Movies[:index], database.Movies[index+1:]...)
			var movie model.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			database.Movies = append(database.Movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode("Movie not found in the database!!")
}

// Delete a movie from the database

func DeleteAMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)

	for index, value := range database.Movies{
		if params["id"] == value.Id{
			database.Movies = append(database.Movies[:index], database.Movies[index+1:]...)
			json.NewEncoder(w).Encode("Movie is deleted from the database")
			return
		}
	}
	json.NewEncoder(w).Encode("Movie not found in the database")
}

// Delete the whole database

func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	database.Movies = []model.Movie{}
	json.NewEncoder(w).Encode("All the movies have been deleted")
	return
}
