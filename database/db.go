package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

// User represents a user document in MongoDB
type User struct {
    Name     string `json:"name" bson:"name"`
    Password string `json:"password" bson:"password"`
}

// GetUsers fetches all users from the "users" collection
func GetUsers() ([]User, error) {
    collection := GetCollection("Strawpoll", "users")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{}) // Empty filter {} gets all users
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var users []User
    if err := cursor.All(ctx, &users); err != nil {
        return nil, err
    }

    return users, nil
}
