// entry point
package main

import (
  // custom
  "github.com/davidlares/books-api/controllers"
  "github.com/davidlares/books-api/models"
  "github.com/davidlares/books-api/driver"
  // env
  "github.com/subosito/gotenv"
  // named
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "database/sql"
)

var books []models.Book

var db *sql.DB

// init
func init() {
  // load all env variables
  gotenv.Load()
}

// main function
func main() {
  // database driver
  db = driver.ConnectDB()
  // creating a router
  router := mux.NewRouter()
  // controller instance
  controller := controllers.Controller{}
  // functions
  router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
  router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
  router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
  router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT") // whole object
  router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")
  // server listening - logging errors
  log.Fatal(http.ListenAndServe(":8000", router))
}
