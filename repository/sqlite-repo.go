package repository

import (
	"database/sql"
	"goapimux/entities"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteRepo struct{}

func NewSQLiteRepository() PostRepository {
	os.Remove("./posts.db")

	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table posts (id integer not null primary key, title text, txt text);
	delete from posts;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}
	return &sqliteRepo{}
}

func (*sqliteRepo) Save(post *entities.Post) (*entities.Post, error) {
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	stmt, err := tx.Prepare("insert into posts(id, title, txt) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(post.ID, post.Title, post.Text)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	tx.Commit()
	return post, nil
}

func (*sqliteRepo) FindAll() ([]entities.Post, error) {
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	rows, err := db.Query("select id, title, txt from posts")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	var posts []entities.Post
	for rows.Next() {
		var id int
		var title string
		var text string
		err = rows.Scan(&id, &title, &text)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		post := entities.Post{
			ID:    id,
			Title: title,
			Text:  text,
		}
		posts = append(posts, post)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return posts, nil
}

func (*sqliteRepo) Delete(post *entities.Post) error {
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}
	stmt, err := tx.Prepare("delete from posts where id = ?")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(post.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	tx.Commit()
	return nil
}
