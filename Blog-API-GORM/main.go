package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func checkerr(err error) {

	if err != nil {
		fmt.Println("Error Happened...", err)
		os.Exit(1)
	}

}

func main() {

	// Loading .env files
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env files")
	}

	PORT := ":" + os.Getenv("PORT")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		//Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Fatal(r.Run(PORT))

}
