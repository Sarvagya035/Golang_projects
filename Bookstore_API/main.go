package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func getallBooks() []Book {

	booksbyte, err := os.ReadFile("./books.json")
	checkerr(err)

	var books []Book

	err = json.Unmarshal(booksbyte, &books)

	return books
}

func saveBook(books []Book) error {

	booksbyte, err := json.Marshal(books)
	checkerr(err)

	err = os.WriteFile("./books.json", booksbyte, 0644)
	checkerr(err)

	return err
}

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Price    string `json:"price"`
	ImageURL string `json:"image_url"`
}

func main() {
	fmt.Println("Server is Starting on Port 4000...")
	r := mux.NewRouter()

	r.HandleFunc("/", ServeHome)
	r.HandleFunc("/getbooks", HandleGetBooks).Methods("GET")
	r.HandleFunc("/book/{id}", HandleGetBookbyID).Methods("GET")
	r.HandleFunc("/add", AddNewBook).Methods("POST")
	r.HandleFunc("/update/{id}", UpdateExistingBook).Methods("PUT")
	r.HandleFunc("/delete/{id}", DeleteExistingBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func checkerr(err error) {

	if err != nil {
		fmt.Println("Error Happened", err)
		os.Exit(1)
	}
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Bookstore API</h1>"))
}

func HandleGetBooks(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get request is Called...")
	w.Header().Set("Content-Type", "applicatioan/json")
	books := getallBooks()
	json.NewEncoder(w).Encode(books)

}

func HandleGetBookbyID(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get Book by ID is called")
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	books := getallBooks()

	var foundbook bool

	for _, book := range books {

		if book.ID == params["id"] {
			foundbook = true
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	if !foundbook {
		fmt.Println("No book found With Given ID ")
		json.NewEncoder(w).Encode("No Book found With given ID")
	}
}

func AddNewBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update book is called...")
	w.Header().Set("Content-type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	books := getallBooks()

	var newbook Book

	_ = json.NewDecoder(r.Body).Decode(&newbook)

	books = append(books, newbook)

	err := saveBook(books)
	checkerr(err)
	fmt.Println("New Book added Sucessfully")

}

func UpdateExistingBook(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Add a new book is called...")
	w.Header().Set("Content-type", "application/json")

	if r.Body == nil {
		fmt.Println("Please Send Some JSON Data")
		return
	}

	books := getallBooks()
	params := mux.Vars(r)

	var updatedBook Book
	var bookFound bool

	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid JSON body")
		return
	}

	for i, book := range books {
		if book.ID == params["id"] {
			updatedBook.ID = book.ID
			books[i] = updatedBook
			bookFound = true
			break

		}
	}

	if !bookFound {
		json.NewEncoder(w).Encode("No book found with given ID")
		return
	}

	err = saveBook(books)
	if err != nil {
		json.NewEncoder(w).Encode("Error while updating book")
		return
	}

	json.NewEncoder(w).Encode("Book updated successfully")

}

func DeleteExistingBook(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Delete Existing book is called...")
	w.Header().Set("Content-type", "application/json")

	books := getallBooks()
	params := mux.Vars(r)

	var foundBook bool

	for i, book := range books {

		if book.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			foundBook = true
			break
		}
	}

	if !foundBook {
		json.NewEncoder(w).Encode("Book with Given ID is not found")
	}

	err := saveBook(books)

	if err != nil {
		json.NewEncoder(w).Encode("Error while updating book")
		return
	}

	json.NewEncoder(w).Encode("Book Deleted Sucessfully")
}
