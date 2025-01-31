package database

import (
	"context"
	"go-api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// GetUsers fetches all users from the "users" collection
func GetUsers() ([]models.User, error) {
    collection := GetCollection("Strawpoll", "users")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{}) // Empty filter {} gets all users
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var users []models.User
    if err := cursor.All(ctx, &users); err != nil {
        return nil, err
    }

    return users, nil
}