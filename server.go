package main

import (
	"goapimux/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.Load()
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", getPosts).Methods("GET")
	router.HandleFunc("/", addPost).Methods("POST")
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}