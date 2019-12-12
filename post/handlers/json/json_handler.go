package json

import (
	"encoding/json"
	"fmt"
	"github.com/omkz/golang-blog/post"
	"net/http"
)

type PostHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type postHandler struct {
	postService post.PostService
}

func NewPostHandler(postService post.PostService) PostHandler {
	return &postHandler{
		postService,
	}
}

func (h *postHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload, err := h.postService.FindAllPosts()
	json.NewEncoder(w).Encode(payload)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

}



