package bookRepository

import (
	"books-list/helpers"
	"books-list/models"
	"database/sql"
)

// BookRepository ...
type BookRepository struct{}

// GetBooks returns all books
func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) []models.Book {
	books = []models.Book{}

	rows, err := db.Query("select * from books")
	helpers.LogFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		helpers.LogFatal(err)
		books = append(books, book)
	}
	return books
}

//GetBook return single book by its id
func (b BookRepository) GetBook(db *sql.DB, id string, book models.Book) models.Book {
	rows := db.QueryRow("select * from books where id=$1", id)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	helpers.LogFatal(err)
	return book
}

// AddBook adds a book
func (b BookRepository) AddBook(db *sql.DB, book models.Book) (bookID int) {
	row := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id",
		book.Title, book.Author, book.Year)
	err := row.Scan(&bookID)

	helpers.LogFatal(err)
	return
}

// UpdateBook updates a book
func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) int64 {
	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
		&book.Title, &book.Author, &book.Year, &book.ID)
	helpers.LogFatal(err)
	rowsUpdated, err := result.RowsAffected()
	helpers.LogFatal(err)
	return rowsUpdated
}

// RemoveBook removes a book by its id
func (b BookRepository) RemoveBook(db *sql.DB, id string) int64 {
	result, err := db.Exec("delete from books where id=$1", id)
	helpers.LogFatal(err)
	rowsUpdated, err := result.RowsAffected()
	helpers.LogFatal(err)
	return rowsUpdated
}
