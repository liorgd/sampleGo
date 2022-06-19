package db

import (
	"database/sql"
	"fmt"
	"main/Structures"
	"strconv"
)

const (
	DB_USER     = "app1"
	DB_PASSWORD = "p"
	DB_NAME     = "app_db"
)

func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)
	return db
}

func GetRecords() []Structures.Record {

	fmt.Println("Getting records...")
	db := SetupDB()

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM records")
	checkErr(err)

	fmt.Println("Parsing records...")
	var records = []Structures.Record{}
	var record = Structures.Record{}
	for rows.Next() {
		var ID string
		var Title string
		var Content string
		var Views int
		var Timestamp int

		err = rows.Scan(&ID, &Title, &Content, &Views, &Timestamp)

		fmt.Println("values: " + ID + " " + Title + " " + Content + " " + strconv.Itoa(Views) + " " + strconv.Itoa(Timestamp))
		record.ID = ID
		record.Title = Title
		record.Content = Content
		record.Views = Views
		record.Timestamp = Timestamp

		records = append(records, record)
		// check errors
		checkErr(err)

	}

	// check errors
	checkErr(err)
	return records
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		fmt.Println("print error")
		panic(err)
	}
}

func AddRecord(newRecord Structures.Record) {

	db := SetupDB()

	fmt.Println("Inserting record into DB")

	var lastInsertID int
	err := db.QueryRow("INSERT INTO records(id, title, content, views, timestamp) VALUES($1, $2, $3, $4, $5) returning id;", newRecord.ID, newRecord.Title, newRecord.Content, newRecord.Views, newRecord.Timestamp).Scan(&lastInsertID)

	// check errors
	checkErr(err)

	fmt.Println("The record has been inserted successfully!")
}
