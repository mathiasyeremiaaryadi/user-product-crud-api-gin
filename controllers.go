package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// User controllers
func GetUsers(c *gin.Context) {
	user_model := []User{}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	if err := db.Limit(limit).Offset(offset).Find(&user_model); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"data":    nil,
			"message": "Failed to retrieve all users",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    user_model,
		"message": "All users retrieved successfully",
	})
}

func GetUser(c *gin.Context) {
	user_model := User{}
	user_id := c.Param("id")

	if err := db.First(&user_model, user_id); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"data":    nil,
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    user_model,
		"message": "A user retrieved sucessfully",
	})
}

func CreateUser(c *gin.Context) {
	var user_model User
	if err := c.ShouldBindJSON(&user_model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"data":    nil,
			"message": "Failed to convert JSON",
		})
		return
	}

	if err := db.Create(&user_model); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"data":    nil,
			"message": "Failed to create new user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"data":    nil,
		"message": "User created successfully",
	})
}

func UpdateUser(c *gin.Context) {
	var user_model User
	user_id := c.Param("id")

	if err := c.ShouldBindJSON(&user_model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"data":    nil,
			"message": "Failed to convert JSON",
		})
		return
	}

	if err := db.Model(&user_model).Where("id=?", user_id).Updates(&user_model); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"data":    nil,
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"data":    nil,
		"message": "User updated successfully",
	})
}

func DeleteUser(c *gin.Context) {
	var user_model User
	user_id := c.Param("id")

	if err := db.Delete(&user_model, user_id); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"data":    nil,
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    nil,
		"message": "User deleted successfully",
	})
}

// Product controllers
func GetProducts(c *gin.Context) {
	product_model := []Product{}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	if err := db.Limit(limit).Offset(offset).Find(&product_model); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"data":    nil,
			"message": "Failed to retrieve all products",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    product_model,
		"message": "All products retrieved successfully",
	})
}

func GetProduct(c *gin.Context) {
	product_model := Product{}
	product_id := c.Param("id")

	if err := db.First(&product_model, product_id); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"data":    nil,
			"message": "Product not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    product_model,
		"message": "A product retrieved sucessfully",
	})
}

func CreateProduct(c *gin.Context) {
	var product_model Product
	if err := c.ShouldBindJSON(&product_model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"data":    nil,
			"message": "Failed to convert JSON",
		})
		return
	}

	if err := db.Create(&product_model); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"data":    nil,
			"message": "Failed to create new product",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"data":    nil,
		"message": "Product created successfully",
	})
}

func UpdateProduct(c *gin.Context) {
	var product_model Product
	product_id := c.Param("id")

	if err := c.ShouldBindJSON(&product_model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"data":    nil,
			"message": "Failed to convert JSON",
		})
		return
	}

	if err := db.Model(&product_model).Where("id=?", product_id).Updates(&product_model); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"data":    nil,
			"message": "Product not found",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"data":    nil,
		"message": "Product updated successfully",
	})
}

func DeleteProduct(c *gin.Context) {
	var product_model Product
	product_id := c.Param("id")

	if err := db.Delete(&product_model, product_id); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"data":    nil,
			"message": "Product not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    nil,
		"message": "Product deleted successfully",
	})
}
