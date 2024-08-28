package controllers

import (
	"context"
	"fmt"
	"log"

	"github.com/mhShohan/go-playground/tree/main/foundation/24-api-mongoDB/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "go-mongo-api"
const collName = "watch_list"

var collection *mongo.Collection

func init() {
	// Set client options
	clientOption := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	checkNil(err)

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)
}

// CheckNil checks if the error is nil
func checkNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// InsertOneMovie inserts a movie into the database
func insertOneMovie(movie model.Netflix) {
	data, err := collection.InsertOne(context.Background(), movie)
	checkNil(err)

	fmt.Println("Inserted a single document: ", data.InsertedID)
}

// Update one watch list
func updateOneWatchList(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	checkNil(err)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"Watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	checkNil(err)

	fmt.Printf("Updated %v document: ", result.ModifiedCount)
}
