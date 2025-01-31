package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// ConnectMongoDB initializes the MongoDB connection
func ConnectMongoDB() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Replace with your MongoDB connection string
    uri := "mongodb://localhost:27017" // For local MongoDB
    // uri := "mongodb+srv://<username>:<password>@cluster0.mongodb.net/?retryWrites=true&w=majority" // For MongoDB Atlas

    clientOptions := options.Client().ApplyURI(uri)

    var err error
    Client, err = mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal("Error connecting to MongoDB:", err)
    }

    // Ping the database
    err = Client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("Could not connect to MongoDB:", err)
    }

    fmt.Println("Connected to MongoDB!")
}

// GetCollection returns a reference to a MongoDB collection
func GetCollection(databaseName, collectionName string) *mongo.Collection {
    return Client.Database(databaseName).Collection(collectionName)
}


