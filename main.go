package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book struct
type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Init books var as a slice Book
var books []Book

// Get All books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get one book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params from request
	// Loop through books and find one with the same id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // MOCK ID, NOT SAFE
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update Book
func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Delete Book
func deleteBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Hardcoded data
	books = append(books, Book{ID: "1", ISBN: "438227", Title: "Book 1", Author: &Author{FirstName: "Rain", LastName: "Smith"}})
	books = append(books, Book{ID: "2", ISBN: "438427", Title: "Book 2", Author: &Author{FirstName: "Sally", LastName: "Mckennie"}})

	// Route handleds & endpoints
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/id={id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBooks).Methods("POST")
	r.HandleFunc("/books/id={id}", updateBooks).Methods("PUT")
	r.HandleFunc("/books/id{id}", deleteBooks).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))

}
