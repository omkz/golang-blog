package post

type PostRepository interface {
	FindAll() ([]*Post, error)
	Create(post *Post) error
	FindById(id string) (*Post, error)
	Delete(id string) error
	Update(id string, post *Post) error
}

