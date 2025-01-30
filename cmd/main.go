package main

import (
	"go-api/routes"
)

func main() {
    router := routes.SetupRouter()
    router.Run("localhost:8080")
}