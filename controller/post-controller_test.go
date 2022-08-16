package controller

import (
	"bytes"
	"encoding/json"
	"goapimux/entities"
	"goapimux/repository"
	"goapimux/services"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	postRepoMock repository.PostRepository = repository.NewSQLiteRepository()
	postServiceMock services.PostService = services.NewPostService(postRepoMock)
	postController PostController = NewPostController(postServiceMock)
)

const (
	TITLE = "title1"
	TEXT = "text1"
)

func TestAddPost(t *testing.T) {
	//Create a new HTTP POST request
	var jsonBody = []byte(`{"title": "` + TITLE + `", "text": "` + TEXT + `"}`)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonBody))

	//Assign HTTP Handler function (controller add post)
	handler := http.HandlerFunc(postController.AddPost)

	//Record HTTP Response (httptest)
	response := httptest.NewRecorder()

	//Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertions on the HTTP Status code and the response
	status := response.Code
	
	//Assert
	assert.Equal(t, http.StatusCreated, status)

	cleanDatabase()
}

func TestGetPosts(t *testing.T) {
	post := entities.Post{
		Title: TITLE,
		Text: TEXT,
	}
	insertRowDatabase(&post)
	req, _ := http.NewRequest("GET", "/", nil)
	handler := http.HandlerFunc(postController.GetPosts)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	status := response.Code
	var posts []entities.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&posts) 
	assert.NotNil(t, posts)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, TITLE, posts[0].Title)
	assert.Equal(t, TEXT, posts[0].Text)

	cleanDatabase()
}

func insertRowDatabase(post *entities.Post) {
	postRepoMock.Save(post)
}

func cleanDatabase() {
	os.Remove("./posts.db")
}