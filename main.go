package main

import (
	"fmt"
	"log"
	"movie_app/router"
	"net/http"
)

func main() {
	fmt.Println("Welcome to my Movie Application!!!")
	r := router.Init()
	log.Fatal(http.ListenAndServe(":2000", r))
}
