package mongodb

import (
	"context"
	"fmt"
	"github.com/omkz/golang-blog/post"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type postRepository struct {
	db *mongo.Client
}

func NewMongoPostRepository(db *mongo.Client) post.PostRepository {
	return &postRepository{
		db,
	}
}

func (r *postRepository) FindAll() (posts []*post.Post, err error) {

	collection := r.db.Database("blog").Collection("posts")

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(),bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	// Here's an array in which you can store the decoded documents
	var results []*post.Post

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		var post post.Post
		err = cur.Decode(&post)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		results = append(results, &post)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results, nil
}

func (r *postRepository) Create(post *post.Post) error {

	collection := r.db.Database("blog").Collection("posts")

	insertResult, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	return  nil
}

func (r *postRepository) FindById(id string) (*post.Post, error) {
	post := new(post.Post)
	filter := bson.D{{"id", id}}
	err := r.db.Database("blog").Collection("posts").FindOne(context.TODO(), filter).Decode(&post)

	if err != nil {
		log.Fatal(err)
	}

	return post, nil
}

