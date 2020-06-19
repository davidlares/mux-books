package controllers

import (
  "github.com/davidlares/books-api/repository/book"
  "github.com/davidlares/books-api/models"
  "github.com/gorilla/mux"
  "encoding/json"
  "database/sql"
  "net/http"
  "strconv"
  "log"
)

type Controller struct{}

// book slice (handling book resources)
var books []models.Book

// logFatal
func logFatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
  // handling functions
  return func (w http.ResponseWriter, r *http.Request) {
    log.Println("Get all books")
    // type Book (struct)
    var book models.Book
    // shelf
    books = []models.Book{}
    // calling repo struct
    bookRepo := bookRepository.BookRepository{}
    // calling method
		books = bookRepo.GetBooks(db, book, books)
    // returning books
    json.NewEncoder(w).Encode(books)
  }
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
  // handling functions
  return func (w http.ResponseWriter, r *http.Request) {
    log.Println("Get one book")
    var book models.Book
    // getting request variables
    params := mux.Vars(r)
    // model
    books = []models.Book{}
    // calling repo struct
    bookRepo := bookRepository.BookRepository{}
    // transforming params (data type)
    id, err := strconv.Atoi(params["id"])
    logFatal(err)
    // calling method
    book = bookRepo.GetBook(db, book, id)
    // return value
    json.NewEncoder(w).Encode(book)
  }
}

// curl -X POST -H "Content-Type: application/json" --data '{"id": 3, "title": "C++ is Old", "author": "Mr C++", "year":"2010"}' http://localhost:8000/books
func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
  // handling functions
  return func (w http.ResponseWriter, r *http.Request) {
    log.Println("Add one book")
    var book models.Book
    var bookID int
    // decoding params
    json.NewDecoder(r.Body).Decode(&book)
    // calling repo struct
    bookRepo := bookRepository.BookRepository{}
    // calling method
    bookID = bookRepo.AddBook(db, book)
    // encoding
    json.NewDecoder(r.Body).Decode(bookID)
  }
}

// curl -X PUT -H "Content-Type: application/json" --data '{"id": 2, "title": "C++ is Great", "author": "Mr C++", "year":"2020"}' http://localhost:8000/books
func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
  // handling functions
  return func (w http.ResponseWriter, r *http.Request) {
    log.Println("Updates a book")
    // type Book (struct)
    var book models.Book
    // map fields to the fields inside the variable book
    json.NewDecoder(r.Body).Decode(&book)
    // calling repo struct
    bookRepo := bookRepository.BookRepository{}
    // update method
    rowsUpdated := bookRepo.UpdateBook(db, book)
    // getting the slice - return
    json.NewEncoder(w).Encode(rowsUpdated)
  }
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
  // handling functions
  return func (w http.ResponseWriter, r *http.Request) {
    log.Println("Delete a book")
    // params map
    params := mux.Vars(r)
    // query
    bookRepo := bookRepository.BookRepository{}
    id, err := strconv.Atoi(params["id"])
    logFatal(err)
    // deleted
    rowsDeleted := bookRepo.RemoveBook(db, id)
    logFatal(err)
    // getting the slice - return
    json.NewEncoder(w).Encode(rowsDeleted)
  }
}
