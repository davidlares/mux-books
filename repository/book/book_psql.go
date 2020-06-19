package bookRepository

import (
  "github.com/davidlares/books-api/models"
  "database/sql"
  "log"
)

type BookRepository struct{}

func logFatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func(b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) []models.Book {
  // get all books
  rows, err := db.Query("SELECT * FROM books")
  logFatal(err)
  // defer close connection
  defer rows.Close()
  // iteration
  for rows.Next() {
    err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
  	logFatal(err)
    books = append(books, book)
  }
  // returning slice
  return books
}

// /GET/{id} - returning a Book objects
func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) models.Book {
  // get one record
	rows := db.QueryRow("SELECT * FROM books WHERE id=$1", id)
  // showing data
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)
  // returning book
	return book
}

// curl -X POST -H "Content-Type: application/json" --data '{"id": 3, "title": "C++ is Old", "author": "Mr C++", "year":"2010"}' http://localhost:8000/books
func (b BookRepository) AddBook(db *sql.DB, book models.Book) int {
  // adding book
  err := db.QueryRow("INSERT INTO books (title, author, year) VALUES($1, $2, $3) RETURNING id;", book.Title, book.Author, book.Year).Scan(&book.ID)
  logFatal(err)
  return book.ID
}

// curl -X PUT -H "Content-Type: application/json" --data '{"id": 2, "title": "C++ is Great", "author": "Mr C++", "year":"2020"}' http://localhost:8000/books
func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) int64 {
	result, err := db.Exec("UPDATE books SET title=$1, author=$2, year=$3 WHERE id=$4 RETURNING id", &book.Title, &book.Author, &book.Year, &book.ID)
	logFatal(err)

	rowsUpdated, err := result.RowsAffected()
	logFatal(err)

	return rowsUpdated
}

// DELETE
func (b BookRepository) RemoveBook(db *sql.DB, id int) int64 {
	result, err := db.Exec("DELETE FROM books WHERE id = $1", id)
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	logFatal(err)

	return rowsDeleted
}
