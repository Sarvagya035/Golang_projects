package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

const PORT = 4000

var My_Secret_Key string = "SARvagya@jwtkey"

type Message struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

// credentials to generate the jwt token
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//custom claim needed for generating jwt token

type MyCustomClaims struct {
	Username     string `json:"username"`
	LoggedinTime string
	jwt.RegisteredClaims
}

func jsonMessageByte(status string, msg string) []byte {
	errMessage := Message{status, msg}
	messageByte, _ := json.Marshal(errMessage)
	return messageByte
}

func CreateJWT() (string, error) {

	currentTime := time.Now().Format("02-01-2006 15:04:05")

	myclaim := MyCustomClaims{
		Username:     "Sarvagya",
		LoggedinTime: currentTime,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Minute)),
			Issuer:    "Sarvagya",
		},
	}

	//Generate token with HS256 algorithm and custom claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myclaim)

	//sign the token with our generated secret key

	signedToken, err := token.SignedString([]byte(My_Secret_Key))

	return signedToken, err

}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	var userData User
	err := json.NewDecoder(r.Body).Decode(&userData)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonMessageByte("Failed", "Bad Request - Failed to parse the payload "))
		return
	}

	log.Printf("username is :=  %v and password is %v", userData.Username, userData.Password)
	//username and password is hardcoded .we can use database here.

	if userData.Username == "Admin" && userData.Password == "Adminpswd" {
		token, _ := CreateJWT()
		w.Write(jsonMessageByte("Success", token))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonMessageByte("Failed", "Invalid Credentials"))
	}

}

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write(jsonMessageByte("Success", "Welcome to authentication in golang using jwt and all"))
}

func main() {
	fmt.Println("Jwt using GO-lang")

	r := mux.NewRouter()

	//setting up routes

	r.HandleFunc("/", serveHome)

	r.HandleFunc("/login", loginHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), r))
}
