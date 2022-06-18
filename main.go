package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"main/db"
	"net/http"
)


// represents data about a record.
type Record struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Content string  `json:"content"`
	Views  int32 `json:"views"`
	Timestamp int32 `json:"timestamp"`
}

var records = []Record{
	{ID: "1", Title: "Blue Train", Content: "John Coltrane", Views: 56, Timestamp: 3452341},
	{ID: "2", Title: "Jeru", Content: "John Coltrane", Views: 57, Timestamp: 3452342},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Content: "John Coltrane", Views: 58, Timestamp: 3452343},
}

func getRecords(c *gin.Context) {
	//db.GetRecords();
	c.IndentedJSON(http.StatusOK, records)
}

func postRecord(c *gin.Context) {
	var newRecord Record

	if err := c.BindJSON(&newRecord); err != nil {
		return
	}

	// Add the new record to the slice.
	records = append(records, newRecord)
	c.IndentedJSON(http.StatusCreated, newRecord)
}

// getRecordByID locates the record whose ID value matches the id
// parameter sent by the client, then returns that record as a response.
func getRecordByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range records {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "record not found"})
}

func main() {
	router := gin.Default()
	router.GET("/records", getRecords)
	router.GET("/records/:id", getRecordByID)
	router.POST("/records", postRecord)

	//db.SetupDB()
	db.GetRecords()
	router.Run("localhost:3000")
}


