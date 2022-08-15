package services

import (
	"errors"
	"goapimux/entities"
	"goapimux/repository"
)

type PostService interface {
	Validate(post *entities.Post) error
	Create(post *entities.Post) (*entities.Post, error)
	FindAll() ([]entities.Post, error)
}

type service struct {}

var repo repository.PostRepository

func NewPostService(rep repository.PostRepository) PostService {
	repo = rep
	return &service{}
}


func (*service) Validate(post *entities.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The post Title is empty")
		return err
	}
	if post.Text == "" {
		err := errors.New("The post Text is empty")
		return err
	}
	return nil
}

func (*service) Create(post *entities.Post) (*entities.Post, error) {
	newPost, err := repo.Save(post)
	return newPost, err
}

func (*service) FindAll() ([]entities.Post, error) {
	posts, err := repo.FindAll()
	return posts, err
}