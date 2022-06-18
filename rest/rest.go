package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/Structures"
	"main/db"
	"net/http"
)

var records = []Structures.Record{
	{ID: "1", Title: "Blue Train", Content: "John Coltrane", Views: 56, Timestamp: 3452341},
	{ID: "2", Title: "Jeru", Content: "John Coltrane", Views: 57, Timestamp: 3452342},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Content: "John Coltrane", Views: 58, Timestamp: 3452343},
}

var UseCache bool

func GetRecords(c *gin.Context) {
	fmt.Println("UseCache:",UseCache)
	if UseCache == true {
		c.IndentedJSON(http.StatusOK, records)
	} else {
		var recordsFromDb = db.GetRecords()
		c.IndentedJSON(http.StatusOK, recordsFromDb)
	}

}

func PostRecord(c *gin.Context) {
	var newRecord Structures.Record

	if err := c.BindJSON(&newRecord); err != nil {
		return
	}

	// Add the new record to the slice.
	records = append(records, newRecord)
	c.IndentedJSON(http.StatusCreated, newRecord)
}

func GetRecordByID(c *gin.Context) {
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
