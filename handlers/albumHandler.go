package handlers

import (
	"net/http"

	"go-api/models"

	"github.com/gin-gonic/gin"
)

// GetAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, models.Albums)
}

// GetAlbumByID locates an album by ID and returns it.
func GetAlbumByID(c *gin.Context) {
    id := c.Param("id")

    for _, a := range models.Albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// PostAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
    var newAlbum models.Album

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    models.Albums = append(models.Albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}