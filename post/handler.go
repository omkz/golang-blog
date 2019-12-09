package post

import (
	"fmt"
)

type PostHandler interface {
	Get()
}

type postHandler struct {
	postService PostService
}

func NewPostHandler(postService PostService) PostHandler {
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
