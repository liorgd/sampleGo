// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type DeleteRecord struct {
	ID string `json:"ID"`
}

type DeletedRecord struct {
	ID string `json:"ID"`
}

type GetRecord struct {
	ID string `json:"ID"`
}

type NewRecord struct {
	ID        string `json:"ID"`
	Title     string `json:"title"`
	Content   string `json:"Content"`
	Views     int    `json:"Views"`
	Timestamp int    `json:"Timestamp"`
}

type Record struct {
	ID        string `json:"ID"`
	Title     string `json:"title"`
	Content   string `json:"Content"`
	Views     int    `json:"Views"`
	Timestamp int    `json:"Timestamp"`
}
