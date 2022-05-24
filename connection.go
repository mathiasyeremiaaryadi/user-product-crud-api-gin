package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Connect() {
	var dbUri string

	if CONFIG["DB_PASS"] == "" {
		dbUri = fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", CONFIG["DB_USER"], CONFIG["DB_HOST"], CONFIG["DB_PORT"], CONFIG["DB_NAME"])
	} else {
		dbUri = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", CONFIG["DB_USER"], CONFIG["DB_PASS"], CONFIG["DB_HOST"], CONFIG["DB_PORT"], CONFIG["DB_NAME"])
	}

	db, err = gorm.Open(mysql.Open(dbUri), &gorm.Config{})

	if err != nil {
		log.Println("Database not connected : ", err)
	} else {
		log.Println("Database connected")
	}

	db.Migrator().DropTable(&User{})
	db.Migrator().DropTable(&Product{})

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{})

	db.Create(&User{
		Username: "Admin",
		Email:    "admin@gmail.com",
		Phone:    "089991234",
		Role:     "admin",
		Status:   false,
		Password: "secret123",
	})
}
