package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"main/rest"
)


func main() {
	rest.UseCache = false
	router := gin.Default()
	router.GET("/records", rest.GetRecords)
	router.GET("/records/:id", rest.GetRecordByID)
	router.POST("/records", rest.PostRecord)

	router.Run("localhost:3000")
}


