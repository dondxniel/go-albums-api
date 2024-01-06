package controllers

import (
	"net/http"

	"github.com/dondxniel/new-albums-api/helpers"
	"github.com/dondxniel/new-albums-api/initializers"
	"github.com/dondxniel/new-albums-api/models"
	"github.com/dondxniel/new-albums-api/structs"
	"github.com/gin-gonic/gin"
)

func CreateAlbum(c *gin.Context) {
	// Get data off request body
	var body struct {
		structs.CreateAlbumRequestBody
	}
	c.Bind(&body);
	album := models.Album{
		Title: body.Title,
		Artist: body.Artist,
		Price: body.Price,
	}
	
	// Create a post	
	result := initializers.DB.Create(&album)

	// Error handling
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, helpers.Response(false, "Error creating album", album, result.Error))
		return
	}
	
	// Return response
	c.JSON(http.StatusCreated, helpers.Response(true, "Album successfully created", album, nil))

}