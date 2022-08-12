package main

import (
	"goapimux/config"
	"goapimux/controller"
	"goapimux/http"
)

var (
	postController controller.PostController = controller.NewPostController()
	router http.Router = http.NewMuxRouter()
)

func main() {
	config.Load()
	const port string = ":8000"
	router.GET("/", postController.GetPosts)
	router.POST("/", postController.AddPost)
	router.SERVE(port)
}