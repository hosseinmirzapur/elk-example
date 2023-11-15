package models

type Post struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
