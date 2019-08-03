package models

type Book struct{
	ISBN int `json:"isbn"`
	Name string `json:"name"`
	Description string `json:"desc"`
	Authors []string `json:"authos"`
	Rating int `json:"rating"`
}