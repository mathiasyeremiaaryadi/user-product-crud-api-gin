package main

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username" `
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Status   *bool  `json:"status"`
	Password string `json:"password"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)
	return
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)
	return
}

type Product struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type JWTClaims struct {
	jwt.StandardClaims
	user User `json:"user"`
}
