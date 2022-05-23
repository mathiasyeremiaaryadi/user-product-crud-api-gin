package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func getJwtToken(c *gin.Context) (string, bool) {
	requestHeader := c.GetHeader("Authorization")
	requestHeaderArr := strings.Split(requestHeader, " ")
	if len(requestHeaderArr) != 2 {
		return "", false
	}

	authRequestType := strings.Trim(requestHeaderArr[0], "\n\r\t")
	if strings.ToLower(authRequestType) != strings.ToLower("Bearer") {
		return "", false
	}

	return strings.Trim(requestHeaderArr[1], "\n\t\r"), true
}

func jwtMiddleware(c *gin.Context) {
	jwtToken, ok := getJwtToken(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"data":    nil,
			"message": "Unauthorized to perform this process",
		})
		return
	}

	authenticatedUser, err := validateJwtToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"data":    nil,
			"message": "Unauthorized to perform this process",
		})
		return
	}

	c.Set("data", authenticatedUser)
	c.Writer.Header().Set("Authorization", "Bearer "+jwtToken)
	c.Next()
}

func APIRoute() {
	router := gin.Default()
	server_url := fmt.Sprint(CONFIG["API_URL"])

	router.Use(CORS())

	router.POST("/login", Login)

	router.Use(jwtMiddleware)
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
