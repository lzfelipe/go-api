package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lzfelipe/go-crud/controllers"
	"github.com/lzfelipe/go-crud/initializers"
)


func init() {
	initializers.Env()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.POST("/user", controllers.CreateUser)
	r.POST("/user/login", controllers.LoginUser)
	r.GET("/user/:id", controllers.ReadUser)

	r.Run()
}