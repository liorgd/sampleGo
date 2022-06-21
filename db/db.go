package db

import (
	"database/sql"
	"fmt"
	"main/Structures"
	"strconv"
)

const (
	DB_HOST     = "10.43.85.202"
	DB_PORT     = "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = "p"
	DB_NAME     = "postgres"
)

func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", DB_USER, DB_PASSWORD, DB_NAME, DB_HOST, DB_PORT)
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

func GetRecord(id string) Structures.Record {

	db := SetupDB()

	fmt.Println("getting record from DB")
	var recordReceived Structures.Record
	strQuery := "SELECT id, title, content, views, timestamp FROM records WHERE id = " + "'" + id + "'"
	err := db.QueryRow(strQuery).Scan(&recordReceived.ID, &recordReceived.Title, &recordReceived.Content, &recordReceived.Views, &recordReceived.Timestamp)

	checkErr(err)

	fmt.Println("The record has been got successfully! recordReceived:", recordReceived)

	return recordReceived
}

func DeleteRecord(id string) {
	db := SetupDB()

	fmt.Println("DeleteRecord record from DB")
	strQuery := "DELETE FROM records WHERE id='" +  id + "'"
	_, err := db.Exec(strQuery)

	checkErr(err)

	fmt.Println("The record successfully deleted ")

}
