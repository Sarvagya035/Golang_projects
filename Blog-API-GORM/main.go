package main

import (
	database "blogAPI_GORM/Database"
	routers "blogAPI_GORM/Routers"
	"fmt"
	"log"
	"net/http"
	"os"

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

	//connecting to database

	database.Connect()

	//Routers and stuff

	router := routers.Router()

	log.Fatal(http.ListenAndServe(PORT, router))
}
