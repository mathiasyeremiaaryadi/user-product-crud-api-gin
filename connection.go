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
	var db_uri string

	if CONFIG["DB_PASS"] == "" {
		db_uri = fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", CONFIG["DB_USER"], CONFIG["DB_HOST"], CONFIG["DB_PORT"], CONFIG["DB_NAME"])
	} else {
		db_uri = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", CONFIG["DB_USER"], CONFIG["DB_PASS"], CONFIG["DB_HOST"], CONFIG["DB_PORT"], CONFIG["DB_NAME"])
	}

	db, err = gorm.Open(mysql.Open(db_uri), &gorm.Config{})

	if err != nil {
		log.Println("Database not connected : ", err)
	} else {
		log.Println("Database connected")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{})
}
