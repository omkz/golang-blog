package post


type PostService interface {
	FindAllPosts() ([]*Post, error)
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
