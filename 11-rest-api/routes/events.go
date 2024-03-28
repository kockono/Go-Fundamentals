package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"res-api.com/apis/models"
	// "res-api.com/apis/utils"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	// Permite retornar un JSON
	context.JSON(http.StatusOK, events)

}

func getEvent(context *gin.Context) {
	// Obtiene el parametro de la URL
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event.ID = 1
	event.UserID = int(userId)

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"status": "Event created", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if event.UserID != int(userId) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to update this event"})
		return
	}

	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateEvent.ID = eventId

	err = updateEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "Event updated successfully!"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		handlerError(context, err)
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		handlerError(context, err)
		return
	}

	// Comprueba si el usuario es el propietario del evento
	if event.UserID != int(userId) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to update this event"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "Event deleted successfully!"})
}

func handlerError(context *gin.Context, err error) {
	context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
}
