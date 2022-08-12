package repository

import (
	"goapimux/db"
	"goapimux/entities"
)

type PostRepository interface {
	Save(post *entities.Post) (*entities.Post, error)
	FindAll() (posts []entities.Post, err error)
}

type repo struct{}

func NewPostRepository() PostRepository {
	return &repo{}
}

func (*repo) Save(post *entities.Post) (*entities.Post, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	result, err := conn.Exec(`INSERT INTO posts (title, text) VALUES ($1, $2) RETURNING id`, post.Title, post.Text)
	if err != nil {
		return nil, err
	}
	lastId, _ := result.LastInsertId()
	post.ID = int(lastId)
	return post, nil
}

func (*repo) FindAll() (posts []entities.Post, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rows, err := conn.Query(`SELECT id, title, text FROM posts`)
	for rows.Next() {
		var post entities.Post
		rows.Scan(&post.ID, &post.Title, &post.Text)
		posts = append(posts, post)
	}
	return
}