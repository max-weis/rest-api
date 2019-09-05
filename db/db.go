package db

import (

	// postgres dialect
	_ "github.com/lib/pq"
	"database/sql"

)

//Book is a model for the demo app
type Book struct {
	ISBN        string `json:"isbn"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Author      string `json:"author"`
	Rating      string `json:"rating"`
}

var (
	getOne    = `SELECT "ISBN","Name","Description","Author","Rating" FROM public.book WHERE "ISBN" = $1`
	getAll    = `SELECT "ISBN","Name","Description","Author","Rating" FROM public.book`
	createOne = `INSERT INTO public.book("ISBN","Name","Description","Author","Rating") VALUES($1,$2,$3,$4,$5) RETURNING "ISBN"`
	updateOne = `UPDATE public.book SET "ISBN"=$1,"Name"=$2, "Description"=$3, "Author"=$4, "Rating"=$5 WHERE "ISBN" = $1 RETURNING "ISBN"`
	deleteOne = `DELETE FROM public.book WHERE "ISBN" = $1`
)

// GetBook queries db for a specifi book
func GetBook(db *sql.DB, bookISBN string) (Book, error) {
	book := Book{}
	var isbn string
	var name string
	var desc string
	var author string
	var rating string

	err := db.QueryRow(getOne, bookISBN).Scan(&isbn, &name, &desc, &author, &rating)
	if err != nil {
		return book, err
	}
	book = Book{ISBN: isbn, Name: name, Description: desc, Author: author, Rating: rating}

	return book, nil
}

//GetAllBooks queries the db for all books
func GetAllBooks(db *sql.DB) ([]Book, error) {
	books := []Book{}
	rows, err := db.Query(getAll)
	defer rows.Close()

	if err != nil {
		return books, err
	}

	for rows.Next() {
		var isbn string
		var name string
		var desc string
		var author string
		var rating string

		err = rows.Scan(&isbn, &name, &desc, &author, &rating)
		if err != nil {
			return books, err
		}
		curBook := Book{ISBN: isbn, Name: name, Description: desc, Author: author, Rating: rating}
		books = append(books, curBook)
	}
	return books, nil
}

//CreateBook inserts book into db
func CreateBook(db *sql.DB, b Book) (string, error) {
	var isbn string
	err := db.QueryRow(
		createOne,
		b.ISBN, b.Name, b.Description, b.Author, b.Rating,
	).Scan(&isbn)

	if err != nil {
		return "", err
	}
	return isbn, nil
}

//UpdateBook updates a book
func UpdateBook(db *sql.DB, b Book) (string, error) {
	var isbn string
	err := db.QueryRow(
		updateOne,
		b.ISBN, b.Name, b.Description, b.Author, b.Rating,
	).Scan(&isbn)

	if err != nil {
		return "", err
	}
	return isbn, nil
}

// DeleteBook deletes a book from the db
func DeleteBook(db *sql.DB, bookISBN string) (string, error) {
	_, err := db.Query(
		deleteOne,
		bookISBN,
	)
	if err != nil {
		return "", err
	}
	return bookISBN, nil
}
