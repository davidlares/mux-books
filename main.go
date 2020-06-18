// entry point
package main

import (
  "log"
  "encoding/json"
  "net/http"
  "reflect"
  "strconv"
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

  // appending book slice
  books = append(books, Book{ID: 1, Title: "Golang pointers", Author: "Mr. GoLang", Year: "2010"},
    Book{ID: 2, Title: "Golang routines", Author: "Mr. GoRoutines", Year: "2010"})

  router.HandleFunc("/books", getBooks).Methods("GET")
  router.HandleFunc("/books/{id}", getBook).Methods("GET")
  router.HandleFunc("/books", addBook).Methods("POST")
  router.HandleFunc("/books", updateBook).Methods("PUT") // whole object
  router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")
  // server listening - logging errors
  log.Fatal(http.ListenAndServe(":8000", router))
}

// handling functions
func getBooks(w http.ResponseWriter, r *http.Request) {
  log.Println("Get all books")
  // encoding status data into JSON
  json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
  log.Println("Get one book")
  // getting request variables
  params := mux.Vars(r)
  // converting data
  i, _ := strconv.Atoi(params["id"])
  // printing data type
  log.Println(reflect.TypeOf(i))
  // looping static data
  for _, book := range books {
    if book.ID == i {
      json.NewEncoder(w).Encode(&book) // passing pointer to book slice
    }
  }
}

// curl -X POST -H "Content-Type: application/json" --data '{"id": 3, "title": "C++ is Old", "author": "Mr C++", "year":"2010"}' http://localhost:8000/books
func addBook(w http.ResponseWriter, r *http.Request) {
  log.Println("Add one book")
  // type Book (struct)
  var book Book
  // getting requests body
  _ = json.NewDecoder(r.Body).Decode(&book)
  // appending to slice
  books = append(books, book)
  // returning the whole list of books
  json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
  log.Println("Updates a book")
  // type Book (struct)
  var book Book
  // map fields to the fields inside the variable book
  json.NewDecoder(r.Body).Decode(&book)
  // looping books
  for i, item := range books {
    if item.ID == book.ID {
      books[i] = book
    }
  }
  // getting the slice - return
  json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
  log.Println("Delete a book")
}
