package routes

import (
	"go-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    router.GET("/albums", handlers.GetAlbums)
    router.GET("/albums/:id", handlers.GetAlbumByID)
    router.POST("/albums", handlers.PostAlbums)

    return router
}