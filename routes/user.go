package routes

import (
	"net/http"

	"github.com/emmanuel/rest_project/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not signup."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Sucessfully created user",
	})

}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	
	context.JSON(http.StatusCreated, gin.H{
		"message": "Login Sucessful",
	})

}
