package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// --------------------------------------------
//
//	model
//
// --------------------------------------------
type Book struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Author      *Author `json:"author"`
	Price       float64 `json:"price"`
	IsPublished bool    `json:"is_published"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// fake DB
var books = []Book{
	{ID: "123423", Title: "Book 1", Author: &Author{Name: "Author 1", Email: "author1@gmail.com"}, Price: 100, IsPublished: true},
	{ID: "14343", Title: "Book 2", Author: &Author{Name: "Author 2", Email: "author2@gmail.com"}, Price: 200, IsPublished: true},
}

// service
type BookService interface {
	GetAll() ([]Book, error)
	GetByID(id int) (Book, error)
	Create(book Book) (Book, error)
	Update(id string, book Book) (Book, error)
}

func (book *Book) IsEmpty() bool {
	return book.Title == "" && book.Author == nil && book.Price == 0 && !book.IsPublished
}

// --------------------------------------------
//
//	controllers
//
// --------------------------------------------
// Serve the home page
func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get: /") // Get: /
	w.Write([]byte("<h1>Hello, World!</h1>"))
}

// get all books
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get: /books") // Get: /books
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get book by id
func getBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get id from url
	params := mux.Vars(r)

	fmt.Printf("Get: /books/%v\n", params["id"]) // Get: /books/1

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode("404! Book not found")
}

// create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post: /books") // Post: /books

	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("400! Please send a request body")
	}

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	// check if the book is empty
	if book.IsEmpty() {
		json.NewEncoder(w).Encode("400! Please send a valid request body")
	}

	// generate a random id
	rand.Seed(time.Now().UnixNano())
	book.ID = strconv.Itoa(rand.Intn(100))

	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

// update a book
func updateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// get id from url
	params := mux.Vars(r)
	fmt.Printf("Put: /books/%v\n", params["id"]) // Put: /books

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	for index, item := range books {
		if item.ID == params["id"] {
			book.ID = item.ID
			books[index] = book
			json.NewEncoder(w).Encode(books)
			return
		}
	}

	json.NewEncoder(w).Encode("404! Book not found")
}

// delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// get id from url
	params := mux.Vars(r)
	fmt.Printf("Delete: /books/%v\n", params["id"]) // Delete: /books

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			json.NewEncoder(w).Encode(books)
			return
		}
	}

	json.NewEncoder(w).Encode("404! Book not found")
}

func main() {
	router := mux.NewRouter()

	fmt.Println("Server is running on port 4000")

	// routes
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBookByID).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	// listen to port 4000
	log.Fatal(http.ListenAndServe(":4000", router))
}
