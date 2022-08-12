package repository

import "goapimux/entities"

type PostRepository interface {
	Save(post *entities.Post) (*entities.Post, error)
	FindAll() (posts []entities.Post, err error)
}