package models

type Book struct {
	ISBN        int    `json:"isbn"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Author      string `json:"autho"`
	Rating      int    `json:"rating"`
}
