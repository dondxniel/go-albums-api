package main

import (
	"github.com/dondxniel/new-albums-api/initializers"
	"github.com/dondxniel/new-albums-api/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.Album{})
}