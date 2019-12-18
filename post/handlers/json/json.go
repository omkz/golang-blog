package json

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/omkz/golang-blog/post"
	"net/http"
)

type PostHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload, _ := h.postService.FindAllPosts()
	json.NewEncoder(w).Encode(payload)
}

func (h *postHandler) Create(w http.ResponseWriter, r *http.Request) {
	var post post.Post
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&post)
	_ = h.postService.CreatePost(&post)

	response, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)

}

func (h *postHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	post, _ := h.postService.FindPostById(id)

	response, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func (h *postHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	_ = h.postService.DeletePost(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
