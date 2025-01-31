package main

import (
	"go-api/database"
	"go-api/routes"
)

func main() {
    database.ConnectMongoDB()
    router := routes.SetupRouter()
    
    router.Run("localhost:8080")
}