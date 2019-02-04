package main

import (
	"books-list/controllers"
	"books-list/driver"
	"books-list/models"
	"database/sql"
	"log"
	"net/http"

	"github.com/subosito/gotenv"

	"github.com/gorilla/mux"
)

var books []models.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.ConnectDB()
	router := mux.NewRouter()
	bookController := controllers.BookController{}

	router.HandleFunc("/books", bookController.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", bookController.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", bookController.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", bookController.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", bookController.RemoveBook(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8002", router))
}
