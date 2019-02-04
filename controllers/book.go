package controllers

import (
	"books-list/models"
	bookRepository "books-list/repository/book"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// BookController ...
type BookController struct{}

var books []models.Book

// GetBooks returns all books
func (c BookController) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		books = []models.Book{}

		bookRepo := bookRepository.BookRepository{}
		books = bookRepo.GetBooks(db, book, books)
		json.NewEncoder(w).Encode(books)
	}

}

// GetBook returns single book by its id
func (c BookController) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		id, _ := mux.Vars(r)["id"]
		bookRepo := bookRepository.BookRepository{}
		book = bookRepo.GetBook(db, id, book)
		json.NewEncoder(w).Encode(book)
	}

}

// AddBook adds a book
func (c BookController) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)
		bookRepo := bookRepository.BookRepository{}
		bookID := bookRepo.AddBook(db, book)
		json.NewEncoder(w).Encode(bookID)

	}
}

// UpdateBook updates a book
func (c BookController) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)

		bookRepo := bookRepository.BookRepository{}
		rowsUpdated := bookRepo.UpdateBook(db, book)
		json.NewEncoder(w).Encode(rowsUpdated)
	}

}

// RemoveBook deletes a book by its id
func (c BookController) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		bookRepo := bookRepository.BookRepository{}
		rowsUpdated := bookRepo.RemoveBook(db, id)
		json.NewEncoder(w).Encode(rowsUpdated)

	}

}
