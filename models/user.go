package models

// User represents a user document in MongoDB
type User struct {
    Name     string `json:"name" bson:"name"`
    Password string `json:"password" bson:"password"`
}