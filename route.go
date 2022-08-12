package main

import (
	"encoding/json"
	"fmt"
	"goapimux/entities"
	"goapimux/repository"
	"log"
	"net/http"
)

var repo repository.PostRepository = repository.NewPostRepository()

func getPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		log.Fatalf(err.Error())
		errorHandler(response, http.StatusInternalServerError, "Error on get posts")
		return
	}
	result, err := json.Marshal(posts)
	if err != nil {
		errorHandler(response, http.StatusInternalServerError, "Error marshalling the posts array")
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write(result)
}

func addPost(response http.ResponseWriter, request *http.Request) {
	var post entities.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		errorHandler(response, http.StatusInternalServerError, "Error decode body")
		return
	}
	_, err = repo.Save(&post)
	if err != nil {
		errorHandler(response, http.StatusInternalServerError, "Error on save a new Post")
		return
	}
	response.WriteHeader(http.StatusCreated)
}

func errorHandler(response http.ResponseWriter, statusCode int, message string) {
	response.WriteHeader(statusCode)
	response.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, message)))
}