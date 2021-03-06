package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/omkz/golang-blog/post"
	"github.com/omkz/golang-blog/post/database/mongodb"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/omkz/golang-blog/post/handlers/json"
	//"github.com/omkz/golang-blog/post/handlers/console"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func postgresConnection(database string) *sql.DB {
	fmt.Println("Connecting to PostgreSQL DB")
	db, err := sql.Open("postgres", database)
	if err != nil {
		log.Fatalf("%s", err)
		panic(err)
	}
	return db
}

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
	// postRepo = mongodb.NewMongoPostRepository(mongoConnection("mongodb://localhost:27017"))
	// postRepo = mongodb.NewMongoPostRepository(mongoConnection("mongodb://root:root@localhost:27017"))
	postRepo = mongodb.NewMongoPostRepository(mongoConnection("mongodb://mongodb:27017"))
	
	//postRepo = postgre.NewPostgresPostRepository(postgresConnection("postgresql://omz@localhost/blog_golang?sslmode=disable"))
	postService := post.NewPostService(postRepo)
	//postHandler := console.NewPostHandler(postService)
	//postHandler.Get()
	postHandler := json.NewPostHandler(postService)

	router := mux.NewRouter()
	router.HandleFunc("/posts", postHandler.Get).Methods("GET")
	router.HandleFunc("/posts", postHandler.Create).Methods("POST")
	router.HandleFunc("/posts/{id}", postHandler.GetById).Methods("GET")
	router.HandleFunc("/posts/{id}/delete", postHandler.Delete).Methods("DELETE")
	router.HandleFunc("/posts/{id}/update", postHandler.Update).Methods("PUT")
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", router))
}