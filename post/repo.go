package post

type PostRepository interface {
	FindAll() ([]*Post, error)
	Create(post *Post) error
}

