package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"res-api.com/apis/models"
)

func registerForEvent(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(400, gin.H{"error": "id is required"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch event."})
		return
	}

	err = event.Register(userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not register for event."})
		return
	}

	context.JSON(200, gin.H{
		"message": "registerForEvent",
	})
}

func cancelRegistration(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(400, gin.H{"error": "id is required"})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not register for event."})
		return
	}

	context.JSON(200, gin.H{
		"message": "registerForEvent",
	})
}
