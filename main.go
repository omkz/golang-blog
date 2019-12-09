package main

import (
	"context"
	"fmt"
	"github.com/omkz/golang-blog/post"
	"github.com/omkz/golang-blog/post/database/mongodb"
	"github.com/omkz/golang-blog/post/presenters/console"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func mongoConnection(uri string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
func main() {
	var postRepo post.PostRepository
	postRepo = mongodb.NewMongoPostRepository(mongoConnection("mongodb://localhost:27017"))
	postService := post.NewPostService(postRepo)
	postHandler := console.NewPostHandler(postService)
	postHandler.Get()
}