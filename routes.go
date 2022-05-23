package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func APIRoute() {
	router := gin.Default()
	server_url := fmt.Sprint(CONFIG["API_URL"])

	router.Use(corsMiddleware())

	router.POST("/login", Login)

	router.Use(jwtMiddleware())
	router.GET("/users", GetUsers)
	router.GET("/users/:id", GetUser)
	router.POST("/users", CreateUser)
	router.PUT("/users/:id", UpdateUser)
	router.DELETE("/users/:id", DeleteUser)

	router.GET("/products", GetProducts)
	router.GET("/products/:id", GetProduct)
	router.POST("/products", CreateProduct)
	router.PUT("/products/:id", UpdateProduct)
	router.DELETE("/products/:id", DeleteProduct)

	router.Run(server_url)
}
