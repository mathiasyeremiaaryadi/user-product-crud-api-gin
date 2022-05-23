package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
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

func jwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken, ok := getJwtToken(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"data":    nil,
				"message": "Unauthorized to perform and access this process, the token is not found",
			})
			return
		}

		authenticatedUser, err := validateJwtToken(jwtToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"data":    nil,
				"message": "Unauthorized to perform and access this process, invalid token",
			})
			return
		}

		c.Set("data", authenticatedUser)
		c.Writer.Header().Set("Authorization", "Bearer "+jwtToken)
		c.Next()
	}

}
