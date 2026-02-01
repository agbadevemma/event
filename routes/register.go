package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/emmanuel/rest_project/models"
	"github.com/gin-gonic/gin"
)

func RegisterforEvents(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("error", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse int"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		fmt.Println("error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch Event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		fmt.Println("error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Register "})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": " Registered "})
}

func CancelRegistrationForEvents(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		fmt.Println("error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not CancelRegistration"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
