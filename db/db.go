package db

import (
	"database/sql"
	"fmt"
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
func printMessage(message string) {
	fmt.Println(message)
}

func GetRecords() {
	db := SetupDB()

	printMessage("Getting records...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM records")
	checkErr(err)

	printMessage("Parsing records...")
	for rows.Next() {
		var s1 string
		var s2 string
		var s3 string
		var i4 string
		var i5 string

		err = rows.Scan(&s1, &s2, &s3, &i4, &i5)

		printMessage("values: " + s1 + " " + s2 + " " + s3 + " " + i4 + " " + i5)

		// check errors
		checkErr(err)

	}

	// check errors
	checkErr(err)

}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		printMessage("print error")
		panic(err)
	}
}
