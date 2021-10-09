package main

import (
	"context"
	"log"
	"net/http"

	"example.com/Handler_file"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func api() {

	mongoClient, err := mongo.Connect(context.Background(), &options.ClientOptions{
		Auth: &options.Credential{
			Username: "udit",
			Password: "udit",
		},
	})
	if err != nil {
		log.Fatalf("Unable to connect to db\n[Error]: \v", err)
	}

	mongoClient.Database("API").CreateCollection(context.Background(), "Users")
	mongoClient.Database("API").CreateCollection(context.Background(), "Posts")

	Collection_users := mongoClient.Database("Insta").Collection("Users")
	Collection_posts := mongoClient.Database("Insta").Collection("Posts")

	Handler_user := Handler_file.NHandler_user(Collection_users)
	Handler_post := Handler_file.NHandler_post(Collection_posts)
	Handler_post_user := Handler_file.NHandler_post_user(Collection_posts)

	http.Handle("/users/", Handler_user)
	http.Handle("/posts/", Handler_post)
	http.Handle("/posts/users/", Handler_post_user)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
