package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book
type Book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Collection of books var
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	json.NewEncoder(w).Encode(Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = id

	for index, item := range books {
		if item.ID == id {
			books = append(books[:index], books[index+1:]...)
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	w.Write([]byte("Can't find a book with the given id"))

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	for index, item := range books {
		if item.ID == id {
			books = append(books[:index], books[index+1:]...)
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.Write([]byte("Can't find a book with the given id"))
}

func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One",
		Author: Author{Firstname: "John", Lastname: "Doe"}})

	books = append(books, Book{ID: "2", Isbn: "847564", Title: "Book Two",
		Author: Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route Handlers
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
