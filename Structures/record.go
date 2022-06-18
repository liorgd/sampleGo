package Structures

type Record struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Content string  `json:"content"`
	Views  int `json:"views"`
	Timestamp int `json:"timestamp"`
}


