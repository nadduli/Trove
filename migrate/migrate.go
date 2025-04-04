package main

import (
	"github.com/nadduli/Trove/initializers"
	"github.com/nadduli/Trove/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
