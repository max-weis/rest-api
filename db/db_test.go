package db

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

var (
	TestGetOne    = `SELECT * FROM public.book WHERE "ISBN" = ?`
	TestGetAll    = `SELECT * FROM public.book`
	TestCreateOne = `INSERT INTO public.book("ISBN","Name","Description","Author","Rating") VALUES($1,$2,$3,$4,$5) RETURNING "ISBN"`
	TestUpdateOne = `UPDATE public.book SET "ISBN"=$1,"Name"=$2, "Description"=$3, "Author"=$4, "Rating"=$5 WHERE "ISBN" = $1 RETURNING "ISBN"`
	TestDeleteOne = `DELETE FROM public.book WHERE "ISBN" = $1`
)

func TestGetBook(t *testing.T) {
	// create mock db and object
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Could not create mock db. err: %s", err)
	}
	defer db.Close()

	// insert fake data into mock object
	row := sqlmock.NewRows([]string{"isbn", "name", "description", "author", "rating"}).
		AddRow("1", "Book", "Book about Books", "Max", "5")

	// tell mock object what to expect as the query
	// ExpectQuery accepts an regex expression not a string
	mock.ExpectQuery(regexp.QuoteMeta(TestGetOne)).
		WithArgs("1").
		WillReturnRows(row)

	book := Book{
		ISBN:        "1",
		Name:        "Book",
		Description: "Book about Books",
		Author:      "Max",
		Rating:      "5",
	}

	// run our function against the fake db
	gotBook, err := GetBook(db, "1")
	if err != nil {
		t.Fatalf("Could not read mock db. err: %s", err)
	}
	assert.Equal(t, gotBook, book)
}

func TestGetAllBooks(t *testing.T) {
	// create mock db and object
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Could not create mock db. err: %s", err)
	}
	defer db.Close()

	// insert fake data into mock object
	rows := sqlmock.NewRows([]string{"isbn", "name", "description", "author", "rating"}).
		AddRow("1", "Book1", "Book about Books", "Max", "5").
		AddRow("2", "Book2", "Book about Books", "Max", "5")

	// tell mock object what to expect as the query
	// ExpectQuery accepts an regex expression not a string
	mock.ExpectQuery(regexp.QuoteMeta(TestGetAll)).
		WillReturnRows(rows)

	books := []Book{
		{
			ISBN:        "1",
			Name:        "Book1",
			Description: "Book about Books",
			Author:      "Max",
			Rating:      "5",
		}, {
			ISBN:        "2",
			Name:        "Book2",
			Description: "Book about Books",
			Author:      "Max",
			Rating:      "5",
		},
	}

	// run out function against the fake db
	gotBook, err := GetAllBooks(db)
	if err != nil {
		t.Fatalf("Could not read mock db. err: %s", err)
	}

	assert.Equal(t, gotBook, books)
}

func TestCreateBook(t *testing.T) {
	// create mock db and object
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Could not create mock db. err: %s", err)
	}
	defer db.Close()

	book := Book{
		ISBN:        "1",
		Name:        "Book",
		Description: "Book about Books",
		Author:      "Max",
		Rating:      "5",
	}

	// tell mock object what to expect as the query
	// ExpectQuery accepts an regex expression not a string
	mock.ExpectQuery(regexp.QuoteMeta(TestCreateOne)).
		//add arguments
		WithArgs(book.ISBN, book.Name, book.Description, book.Author, book.Rating).
		// because the query returns the isbn of the book, there needs to be a table with the wanted isbn
		WillReturnRows(sqlmock.NewRows([]string{"ISBN"}).AddRow(book.ISBN))

	// run out function against the fake db
	isbn, err := CreateBook(db, book)
	if err != nil {
		t.Fatalf("Could not read mock db. err: %s", err)
	}

	assert.Equal(t, isbn, book.ISBN)
}

func TestUpdateBook(t *testing.T) {
	// create mock db and object
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Could not create mock db. err: %s", err)
	}
	defer db.Close()

	// insert fake data into mock object
	_ = sqlmock.NewRows([]string{"isbn", "name", "description", "author", "rating"}).
		AddRow("1", "Book", "Book about Books", "Max", "5")

	// the new updates (rating isn't that good :D)
	book := Book{
		ISBN:        "1",
		Name:        "Book",
		Description: "Book about Books",
		Author:      "Max",
		Rating:      "1",
	}

	// tell mock object what to expect as the query
	// ExpectQuery accepts an regex expression not a string
	mock.ExpectQuery(regexp.QuoteMeta(TestUpdateOne)).
		//add arguments
		WithArgs(book.ISBN, book.Name, book.Description, book.Author, book.Rating).
		// because the query returns he isbn of the book, there needs to be a table with the wanted isbn
		WillReturnRows(sqlmock.NewRows([]string{"ISBN"}).AddRow(book.ISBN))

	// run out function against the fake db
	isbn, err := UpdateBook(db, book)
	if err != nil {
		t.Fatalf("Could not read mock db. err: %s", err)
	}

	assert.Equal(t, isbn, book.ISBN)
}

func TestDeleteBook(t *testing.T) {
	// create mock db and object
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Could not create mock db. err: %s", err)
	}
	defer db.Close()

	// insert fake data into mock object
	_ = sqlmock.NewRows([]string{"isbn", "name", "description", "author", "rating"}).
		AddRow("1", "Book", "Book about Books", "Max", "5")

	// tell mock object what to expect as the query
	// ExpectQuery accepts an regex expression not a string
	mock.ExpectQuery(regexp.QuoteMeta(TestDeleteOne)).
		//add arguments
		WithArgs("1").
		// because the query returns he isbn of the book, there needs to be a table with the wanted isbn
		WillReturnRows(sqlmock.NewRows([]string{"ISBN"}).AddRow(1))

	// run out function against the fake db
	isbn, err := DeleteBook(db, "1")
	if err != nil {
		t.Fatalf("Could not read mock db. err: %s", err)
	}

	assert.Equal(t, isbn, "1")
}
