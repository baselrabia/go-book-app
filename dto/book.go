package dto

type Book struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Published int    `json:"published"`
}