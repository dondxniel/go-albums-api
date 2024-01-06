package controllers

import (
	"net/http"

	"github.com/dondxniel/new-albums-api/helpers"
	"github.com/dondxniel/new-albums-api/initializers"
	"github.com/dondxniel/new-albums-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteAlbum(c *gin.Context) {
	// Get ID from url
	id := c.Param("id")

	// Delete the item
	result := initializers.DB.Delete(&models.Album{}, id)

	// Error handling
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, helpers.Response(false, "Album not found", nil, result.Error))
			return
		}
		c.JSON(http.StatusBadRequest, helpers.Response(false, "Error deleting album", nil, result.Error))
		return
	}

	c.JSON(http.StatusOK, helpers.Response(false, "Item successfully deleted", nil, nil))
}