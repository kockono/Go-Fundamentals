package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"res-api.com/apis/db"
	"res-api.com/apis/models"
)

func main() {
	db.InitDb()
	server := gin.Default() // create a new server

	server.GET("/events", getEvents) // GET, POST, PUT, DELETE
	server.POST("/events", createEvent)
	server.Run(":8080") // localhost:8080

}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	// Permite retornar un JSON
	context.JSON(http.StatusOK, events)
	// Permite retornar un HTML
	// context.HTML(http.StatusOK, "<h1>Events</h1>", nil)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return // Exit the function
	}

	event.ID = 1
	event.UserID = 1
	event.Save()
	events := models.GetAllEvents()
	context.JSON(http.StatusCreated, gin.H{"status": "Event created", "event": events})
}
