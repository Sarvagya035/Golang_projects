package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const PORT = 4000

type Message struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func jsonMessageByte(status string, msg string) []byte {
	errMessage := Message{status, msg}
	messageByte, _ := json.Marshal(errMessage)
	return messageByte
}

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write(jsonMessageByte("Success", "Welcome to authentication in golang using jwt and all"))
}

func main() {
	fmt.Println("Jwt using GO-lang")

	r := mux.NewRouter()

	//setting up routes

	r.HandleFunc("/", serveHome)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), r))
}
