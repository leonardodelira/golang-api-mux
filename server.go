package main

import (
	"goapimux/config"
	"goapimux/controller"
	"goapimux/http"
	"goapimux/repository"
	"goapimux/services"
)

var (
	postRepo repository.PostRepository = repository.NewPostgresRepository()
	postService services.PostService = services.NewPostService(postRepo)
	postController controller.PostController = controller.NewPostController(postService)
	router http.Router = http.NewMuxRouter()
)

func main() {
	config.Load()
	const port string = ":8000"
	router.GET("/", postController.GetPosts)
	router.POST("/", postController.AddPost)
	router.SERVE(port)
}