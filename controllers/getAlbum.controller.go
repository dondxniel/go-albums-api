package controllers

import (
	"net/http"

	"github.com/dondxniel/new-albums-api/helpers"
	"github.com/dondxniel/new-albums-api/initializers"
	"github.com/dondxniel/new-albums-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAlbum(c *gin.Context){
	// Get ID
	id := c.Param("id")
	
	var album models.Album;
	result := initializers.DB.First(&album, id);

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, helpers.Response(false, "Album not found", nil, result.Error))
			return;
		}
		c.JSON(http.StatusBadRequest, helpers.Response(false, "Error getting album from database.", nil, result.Error))
		return;
	}
	
	c.JSON(http.StatusOK, helpers.Response(true, "Album successfully fetched.", album, nil))
}