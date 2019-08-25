package app

import (
	// use a dot import to avoid models.Book
	. "gitlab.com/baroprime/prod-rest/models"

	// postgres dialect
	_ "github.com/lib/pq"
)

var (
	getOne    = `SELECT "ISBN","Name","Description","Author","Rating" FROM public.book WHERE "ISBN" = $1`
	getAll    = `SELECT "ISBN","Name","Description","Author","Rating" FROM public.book`
	createOne = `INSERT INTO public.book("ISBN","Name","Description","Author","Rating") VALUES($1,$2,$3,$4,$5) RETURNING "ISBN"`
	updateOne = `UPDATE public.book SET "ISBN"=$1,"Name"=$2, "Description"=$3, "Author"=$4, "Rating"=$5 WHERE "ISBN" = $1 RETURNING "ISBN"`
	deleteOne = `DELETE FROM public.book WHERE "ISBN" = $1`
)

// GetBook queries db for a specifi book
func GetBook(a *App, bookISBN string) (Book, error) {
	book := Book{}
	var isbn string
	var name string
	var desc string
	var author string
	var rating string

	err := a.DB.QueryRow(getOne, bookISBN).Scan(&isbn, &name, &desc, &author, &rating)
	if err != nil {
		return book, err
	}
	book = Book{ISBN: isbn, Name: name, Description: desc, Author: author, Rating: rating}

	return book, nil
}

//GetAllBooks queries the db for all books
func GetAllBooks(a *App) ([]Book, error) {
	books := []Book{}
	rows, err := a.DB.Query(getAll)
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
func CreateBook(a *App, b Book) (string, error) {
	var isbn string
	err := a.DB.QueryRow(
		createOne,
		b.ISBN, b.Name, b.Description, b.Author, b.Rating,
	).Scan(&isbn)

	if err != nil {
		return "", err
	}
	return isbn, nil
}

//UpdateBook updates a book
func UpdateBook(a *App, b Book) (string, error) {
	var isbn string
	err := a.DB.QueryRow(
		updateOne,
		b.ISBN, b.Name, b.Description, b.Author, b.Rating,
	).Scan(&isbn)

	if err != nil {
		return "", err
	}
	return isbn, nil
}

// DeleteBook deletes a book from the db
func DeleteBook(a *App, bookISBN string) (string, error) {
	_, err := a.DB.Query(
		deleteOne,
		bookISBN,
	)
	if err != nil {
		return "", err
	}
	return bookISBN, nil
}
