package controller

import (
	"encoding/json"
	"goapimux/entities"
	"goapimux/errors"
	"goapimux/services"
	"net/http"
)

var postService services.PostService = services.NewPostService()

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
}

func GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		errorHandler(response, http.StatusInternalServerError, "Error marshalling the posts array")
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func AddPost(response http.ResponseWriter, request *http.Request) {
	var post entities.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		errorHandler(response, http.StatusInternalServerError, "Error decode body")
		return
	}
	err = postService.Validate(&post)
	if err != nil {
		errorHandler(response, http.StatusBadRequest, err.Error())
		return
	}
	postService.Create(&post)
	response.WriteHeader(http.StatusCreated)
}

func errorHandler(response http.ResponseWriter, statusCode int, message string) {
	response.WriteHeader(statusCode)
	json.NewEncoder(response).Encode(errors.ServiceError{Message: message})
}