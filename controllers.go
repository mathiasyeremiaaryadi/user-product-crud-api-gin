package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Authentication controllers
func Login(c *gin.Context) {
	var userModel User

	if err := c.ShouldBindJSON(&userModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"data":    nil,
			"message": "Failed to convert JSON",
		})
		return
	}

	inputEmail := userModel.Email
	inputPassword := []byte(userModel.Password)

	if err := db.Where("email=?", inputEmail).Find(&userModel); err.Error != nil || userModel.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"data":    nil,
			"message": "Invalid email",
		})
		return
	}

	dbUserPassword := []byte(userModel.Password)
	if err := bcrypt.CompareHashAndPassword(dbUserPassword, inputPassword); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"data":    nil,
			"message": "Invalid password",
		})
		return
	}

	jwtToken, err := generateJwtToken(userModel)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"data":    nil,
			"message": "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    jwtToken,
		"message": "Login successfully",
	})
}

// User controllers
func GetUsers(c *gin.Context) {
	userModel := []User{}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	if err := db.Limit(limit).Offset(offset).Find(&userModel); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"data":    nil,
			"message": "Failed to retrieve all users",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    userModel,
		"message": "All users retrieved successfully",
	})
}

func GetUser(c *gin.Context) {
	userModel := User{}
	userId := c.Param("id")

	if err := db.First(&userModel, userId); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"data":    nil,
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    userModel,
		"message": "A user retrieved sucessfully",
	})
}

func CreateUser(c *gin.Context) {
	var userModel User
	if err := c.ShouldBindJSON(&userModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"data":    nil,
			"message": "Failed to convert JSON",
		})
		return
	}

	if err := db.Create(&userModel); err.Error != nil {
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
	var userModel User
	userId := c.Param("id")

	if err := c.ShouldBindJSON(&userModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"data":    nil,
			"message": "Failed to convert JSON",
		})
		return
	}

	if err := db.Model(&userModel).Where("id=?", userId).Updates(&userModel); err.Error != nil {
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
	var userModel User
	userId := c.Param("id")

	if err := db.Delete(&userModel, userId); err.Error != nil {
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
	productModel := []Product{}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	if err := db.Limit(limit).Offset(offset).Find(&productModel); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"data":    nil,
			"message": "Failed to retrieve all products",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    productModel,
		"message": "All products retrieved successfully",
	})
}

func GetProduct(c *gin.Context) {
	productModel := Product{}
	productId := c.Param("id")

	if err := db.First(&productModel, productId); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"data":    nil,
			"message": "Product not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    productModel,
		"message": "A product retrieved sucessfully",
	})
}

func CreateProduct(c *gin.Context) {
	var productModel Product
	if err := c.ShouldBindJSON(&productModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"data":    nil,
			"message": "Failed to convert JSON",
		})
		return
	}

	if err := db.Create(&productModel); err.Error != nil {
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
	var productModel Product
	productId := c.Param("id")

	if err := c.ShouldBindJSON(&productModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"data":    nil,
			"message": "Failed to convert JSON",
		})
		return
	}

	if err := db.Model(&productModel).Where("id=?", productId).Updates(&productModel); err.Error != nil {
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
	var productModel Product
	productId := c.Param("id")

	if err := db.Delete(&productModel, productId); err.Error != nil {
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
