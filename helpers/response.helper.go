package helpers

import "github.com/gin-gonic/gin"

func Response(success bool, message string, data any, error error) *gin.H {
	return &gin.H{
		"success": success,
		"message": message,
		"data": data,
		"error": error,
	}
}