package postgre

import (
	"database/sql"
	"fmt"
	"github.com/omkz/golang-blog/post"
	"log"

	_ "github.com/lib/pq"
)

type postRepository struct {
	db *sql.DB
}

func NewPostgresPostRepository(db *sql.DB) post.PostRepository {
	return &postRepository{
		db,
	}
}

func (r *postRepository) FindAll() (posts []*post.Post, err error) {
	rows, err := r.db.Query("SELECT * FROM posts")
	defer rows.Close()

	for rows.Next() {
		post := new(post.Post)
		if err = rows.Scan(&post.Id, &post.Title, &post.Description, &post.Content, &post.Created_at, &post.Updated_at); err != nil {
			log.Print(err)
			return nil, err
		}
		posts = append(posts, post)

	}

	return posts, nil
}

//func (r *postRepository) Create(post *post.Post) error {
//	_, err := r.db.Exec("INSERT INTO posts(title, content, description, created_at, updated_at) "+
//		"VALUES ($1, $2, $3, $4, $5) RETURNING id",
//		post.Title, post.Description, post.Content, post.Created_at, post.Updated_at)
//
//	if err != nil {
//		fmt.Println(err.Error())
//		return nil
//	}
//
//	return nil
//}

func (r *postRepository) Create(post *post.Post) error {
	 _, err := r.db.Exec("INSERT INTO posts(id, title, content, description, created_at, updated_at) "+
		"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		post.Id, post.Title, post.Description, post.Content, post.Created_at, post.Updated_at)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return nil
}
