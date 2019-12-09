package console

import (
	"fmt"
	"github.com/omkz/golang-blog/post"
)

type PostHandler interface {
	Get()
}

type postHandler struct {
	postService post.PostService
}

func NewPostHandler(postService post.PostService) PostHandler {
	return &postHandler{
		postService,
	}
}

func (h *postHandler) Get() {
	posts, err := h.postService.FindAllPosts()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range posts {
		fmt.Println(each.Title)
		fmt.Println(each.Description)
		fmt.Println(each.Content)
	}
}
