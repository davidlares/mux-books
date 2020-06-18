// entry point
package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

// book model
type Book struct {
  ID int `json:id`
  Title string `json:title`
  Author string `json:author`
  Year string `json:year`
}

// book slice (handling book resources)
var books []Book

// main function
func main() {
  // creating a router
  router := mux.NewRouter()
  router.HandleFunc("/books", getBooks).Methods("GET")
  router.HandleFunc("/books/{id}", getBook).Methods("GET")
  router.HandleFunc("/books", addBook).Methods("POST")
  router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
  router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")
  // server listening - logging errors
  log.Fatal(http.ListenAndServe(":8000", router))
}

// handling functions
func getBooks(w http.ResponseWriter, r *http.Request) {
  log.Println("Get all books")
}

func getBook(w http.ResponseWriter, r *http.Request) {
  log.Println("Get one book")
}

func addBook(w http.ResponseWriter, r *http.Request) {
  log.Println("Add one book")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
  log.Println("Updates a book")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
  log.Println("Delete a book")
}
