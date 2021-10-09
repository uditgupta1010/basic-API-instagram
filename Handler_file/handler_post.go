package Handler_file

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"example.com/Data_file"
	"example.com/Feature_file"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *Handler_post) createPost(w http.ResponseWriter, r *http.Request) {
	post := &Data_file.Post{}
	ok := Feature_file.ReadJson(w, r, post)
	if !ok {
		return
	}

	if err := Feature_file.ValidatePost(post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rand.Seed(time.Now().UnixNano())
	post.Id = strconv.FormatInt(int64(rand.Uint64()), 10)
	post.PostedTimestamp = time.Now()
	_, err := h.Collection_posts.InsertOne(context.Background(), post)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("successfully created post"))
}

func (h *Handler_post) getPost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/posts/"):]
	fmt.Println(id)

	post := &Data_file.Post_j{}
	err := h.Collection_posts.FindOne(context.Background(), bson.D{{"_id", id}}).Decode(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Feature_file.WriteJson(w, r, post)
}

type Handler_post struct {
	Collection_posts *mongo.Collection
}

func NHandler_post(col *mongo.Collection) *Handler_post {
	return &Handler_post{
		Collection_posts: col,
	}
}

func (h *Handler_post) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getPost(w, r)
	case http.MethodPost:
		h.createPost(w, r)
	default:
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}
