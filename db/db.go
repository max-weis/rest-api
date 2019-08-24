package db

import (
	. "gitlab.com/baroprime/prod-rest/app"
	. "gitlab.com/baroprime/prod-rest/models"

	// postgres dialect
	_ "github.com/lib/pq"
)

func GetAllBooks(a *App) ([]Book, error) {
	books := []Book{}
	rows, err := a.DB.Query(`SELECT isbn, name,desc,author,rating FROM books order by isbn`)
	defer rows.Close()

	if err != nil {
		return books, err
	}

	for rows.Next() {
		var isbn int
		var name string
		var desc string
		var author string
		var rating int

		err = rows.Scan(&isbn, &name, &desc, &author, &rating)
		if err != nil {
			return books, err
		}
		curBook := Book{ISBN: isbn, Name: name, Description: desc, Author: author, Rating: rating}
		books = append(books, curBook)
	}
	return books, nil
}

func CreateBook(a *App, b Book) (int, error) {
	var isbn int
	err := a.DB.QueryRow(
		`INSERT INTO books(isbn,name,desc,author,rating) VALUES($1,$2,$3,$4,$5) RETURNING isbn`,
		b.ISBN, b.Name, b.Description, b.Author, b.Rating,
	).Scan(&isbn)

	if err != nil {
		return -1, err
	}
	return isbn, nil
}
