package main

import (
	endpoints "github.com/dondxniel/new-albums-api/constants"
	"github.com/dondxniel/new-albums-api/controllers"
	"github.com/dondxniel/new-albums-api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	router := gin.Default();

	router.POST(endpoints.Base, controllers.CreateAlbum)
	router.GET(endpoints.Base, controllers.GetAlbums)
	router.GET(endpoints.BaseWithId, controllers.GetAlbum)
	router.PATCH(endpoints.BaseWithId, controllers.UpdateAlbum)
	router.DELETE(endpoints.BaseWithId, controllers.DeleteAlbum)

	router.Run()
}