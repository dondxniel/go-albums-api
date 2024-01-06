package controllers

import (
	"net/http"

	"github.com/dondxniel/new-albums-api/helpers"
	"github.com/dondxniel/new-albums-api/initializers"
	"github.com/dondxniel/new-albums-api/models"
	"github.com/dondxniel/new-albums-api/structs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateAlbum(c *gin.Context){
	// Get ID
	id := c.Param("id")

	// Get Data off request body
	var body struct {
		structs.CreateAlbumRequestBody
	}
	c.Bind(&body);

	// Find data 
	var album models.Album;
	result := initializers.DB.First(&album, id);

	// Handle error
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, helpers.Response(false, "Album not found", nil, result.Error))
			return;
		}
		c.JSON(http.StatusBadRequest, helpers.Response(false, "Error getting album from database.", nil, result.Error))
		return;
	}

	// Update
	updateResult := initializers.DB.Model(&album).Updates(models.Album{
		Title: body.Title,
		Artist: body.Artist,
		Price: body.Price,
	})

	if updateResult.Error != nil {
		c.JSON(http.StatusBadRequest, helpers.Response(false, "Error updating fields.", nil, result.Error))
		return;
	}

	// Return 
	c.JSON(http.StatusOK, helpers.Response(true, "Album updated.", album, nil))
}