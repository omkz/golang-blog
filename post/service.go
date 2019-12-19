package post

import (
	"time"
	"github.com/google/uuid"
)


type PostService interface {
	FindAllPosts() ([]*Post, error)
	CreatePost(post *Post) error
	FindPostById(id string) (*Post, error)
	DeletePost(id string) error
	UpdatePost(id string, post *Post) error
}

type postService struct {
	repo PostRepository
}

func NewPostService(repo PostRepository) PostService {
	return &postService{
		repo,
	}
}

func (s *postService) FindAllPosts() ([]*Post, error) {
	return s.repo.FindAll()
}

func(s *postService) CreatePost(post *Post) error{
	post.Id = uuid.New().String()
	post.Created_at = time.Now()
	post.Updated_at = time.Now()
	return s.repo.Create(post)
}

func (s *postService) FindPostById(id string) (*Post, error){
	return s.repo.FindById(id)
}

func (s *postService) DeletePost(id string) error{
	return s.repo.Delete(id)
}

func (s *postService) UpdatePost(id string, post *Post) error{
	post.Updated_at = time.Now()
	return s.repo.Update(id, post)
}

