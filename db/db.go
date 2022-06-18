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

// Function for handling messages
func PrintMessage(message string) {
	fmt.Println(message)
}

func GetRecords() []Structures.Record {

	PrintMessage("Getting records...")
	db := SetupDB()

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM records")
	checkErr(err)

	PrintMessage("Parsing records...")
	var records = []Structures.Record{}
	var record = Structures.Record{}
	for rows.Next() {
		var ID string
		var Title string
		var Content string
		var Views int
		var Timestamp int

		err = rows.Scan(&ID, &Title, &Content, &Views, &Timestamp)

		PrintMessage("values: " + ID + " " + Title + " " + Content + " " + strconv.Itoa(Views) + " " + strconv.Itoa(Timestamp))
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
		PrintMessage("print error")
		panic(err)
	}
}
