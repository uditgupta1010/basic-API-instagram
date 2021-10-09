package Handler_file

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"

	"example.com/Data_file"
	"example.com/Feature_file"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *Handler_user) createUser(w http.ResponseWriter, r *http.Request) {
	user := &Data_file.User{}
	ok := Feature_file.ReadJson(w, r, user)
	if !ok {
		return
	}

	err1 := Feature_file.ValidateUser(user)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword := sha256.New()
	hashedPassword.Write([]byte(user.Password))
	user.Password = fmt.Sprintf("%x\n", hashedPassword.Sum(nil))

	userResult, err := h.Collection_users.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(fmt.Sprintf("Successfully created user with id: %v", userResult.InsertedID)))
}

func (h *Handler_user) getUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/users/"):]
	fmt.Println(id)

	user := &Data_file.User_j{}
	userResult := h.Collection_users.FindOne(context.Background(), bson.D{{"_id", id}})
	err := userResult.Decode(user)
	if err != nil {
		w.Write([]byte("unable to get data"))
	} else {
		Feature_file.WriteJson(w, r, user)
	}

}

type Handler_user struct {
	Collection_users *mongo.Collection
}

func NHandler_user(col *mongo.Collection) *Handler_user {
	return &Handler_user{
		Collection_users: col,
	}
}

func (h *Handler_user) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			h.getUser(w, r)
		}
	case http.MethodPost:
		{
			h.createUser(w, r)
		}
	default:
		{
			http.Error(w, "Method not implemented", http.StatusMethodNotAllowed)
		}
	}
}
