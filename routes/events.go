package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/emmanuel/rest_project/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Println("error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}
	context.JSON(
		http.StatusOK,
		gin.H{
			"events":  events,
			"message": "Success",
		},
	)
}

func getEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse int"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse int"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Success", "event": event})
}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		fmt.Println("err", err)
		return
	}

	userId := context.GetInt64("userId")
	event.ID = userId
	event.UserID = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create parse"})
		fmt.Println("err", err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Print()
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse int"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse int"})
		return
	}
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to update"})
		return
	}

	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "An error occured"})
		fmt.Println("err", err)
		return
	}

	updateEvent.ID = eventId
	err = updateEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create parse"})
		fmt.Println("err", err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Updated"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse int"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not be found"})
		return
	}
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to delete"})
		return
	}

	var deleteEvent models.Event
	deleteEvent.ID = eventId
	err = deleteEvent.DeleteEvent()
	fmt.Println("err", err)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "couldn't delete"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Deleted"})
}
