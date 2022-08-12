package main

import (
	"goapimux/config"
	"goapimux/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.Load()
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", controller.GetPosts).Methods("GET")
	router.HandleFunc("/", controller.AddPost).Methods("POST")
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}