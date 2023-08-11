package main

import (
	"github.com/lzfelipe/go-crud/initializers"
	model "github.com/lzfelipe/go-crud/models"
)

func init() {
	initializers.Env()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&model.User{})
}