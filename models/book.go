package models

//Book is a model for the demo app
type Book struct {
	ISBN        string `json:"isbn"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Author      string `json:"author"`
	Rating      string `json:"rating"`
}
