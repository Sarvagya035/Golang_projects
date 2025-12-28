package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Server is starting on port 4000....")

	r := mux.NewRouter()

	r.HandleFunc("/", serveHome)
	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Welcome to System meterics API "))
}
