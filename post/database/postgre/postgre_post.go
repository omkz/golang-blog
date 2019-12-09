package postgre

import (
	"database/sql"
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
		if err = rows.Scan(&post.Id, &post.Title, &post.Content); err != nil {
			log.Print(err)
			return nil, err
		}
		posts = append(posts, post)

	}

	return posts, nil
}