package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcelomoresco/go-jwt/controller"
	"github.com/marcelomoresco/go-jwt/initializers"
	"github.com/marcelomoresco/go-jwt/middlewares"
)

func init(){
	initializers.LoadEnv()
	initializers.ConnectDb()
}

func main(){
	r := gin.Default()
	r.POST("/signup", controller.Signup)
	r.POST("/login", controller.Login)
	r.GET("/something",middlewares.RequiredAuth, controller.Validate)
	r.Run()
}