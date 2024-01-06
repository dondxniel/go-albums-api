package controllers

import (
	"net/http"

	"github.com/dondxniel/new-albums-api/helpers"
	"github.com/dondxniel/new-albums-api/initializers"
	"github.com/dondxniel/new-albums-api/models"
	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context){
	var albums []models.Album;
	result := initializers.DB.Find(&albums);

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, helpers.Response(false, "Error getting albums from database.", []models.Album{}, result.Error))
		return;
	}
	
	c.JSON(http.StatusOK, helpers.Response(true, "Albums successfully fetched.", albums, nil))
}