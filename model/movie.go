package model

type Movie struct{
	Id string `json:"id"`
	Movie_name string `json:"name"`
	Movie_genre string `json:"genre"`
	Release_year int `json:"year"`
}