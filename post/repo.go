package post

type PostRepository interface {
	FindAll() ([]*Post, error)
}

