package database

import (
	models "blogAPI_GORM/Models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "root:Learningisfun@2026@tcp(127.0.0.1:3306)/Blog_API?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error Connecting to database...")
	}

	DB = db

	fmt.Println("Database Connected Sucessfully")
	db.AutoMigrate(&models.Blog{})
}
