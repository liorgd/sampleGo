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
	fmt.Println("UseCache:", UseCache)
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

	if UseCache == true {
		// Add the new record to the slice.
		records = append(records, newRecord)
		c.IndentedJSON(http.StatusCreated, newRecord)
	} else {
		db.AddRecord(newRecord)
	}
}

func GetRecordByID(c *gin.Context) {
	id := c.Param("id")

	if UseCache == true {
		for _, gotrecord := range records {
			if gotrecord.ID == id {
				c.IndentedJSON(http.StatusOK, gotrecord)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "record not found"})
	} else {
		var recordReceived = db.GetRecord(id)
		c.IndentedJSON(http.StatusOK, recordReceived)
		return
	}
}

func DeleteRecord(c *gin.Context) {
	id := c.Param("id")
	if UseCache == true {
		// not implemented
	} else {
		db.DeleteRecord(id)
		c.IndentedJSON(http.StatusOK, "")
		return
	}

}