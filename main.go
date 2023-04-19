package main

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

// define some package level variables
var (
	collection *mongo.Collection

	// like a timeout or deadline that indicates when an operation
	// should stop running and return. It helps to prevent performance
	// degradation on production sytems.
	ctx context.TODO
)

// a task manager system written in go and using mongoDb
func main() {

}

func init() {

	// options.ClientOption used to set the connection string
	// and other driver settings.
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// let's ensure that out MongoDB server was found 
	// and connected successfully using Ping method.
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}


	// The Database and Collection types can be 
	// used to access the database
	collection = client.Database("task-manager").Collection("tasks")
}