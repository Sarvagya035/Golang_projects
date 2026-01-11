package database

import (
	models "blogAPI_GORM/Models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "PASTE_YOUR_SECRET_KEY_HERE"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error Connecting to database...")
	}

	DB = db

	fmt.Println("Database Connected Sucessfully")
	db.AutoMigrate(&models.Blog{})
}
