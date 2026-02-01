package routes

import (
	"net/http"

	"github.com/emmanuel/rest_project/models"
	"github.com/emmanuel/rest_project/utils"
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

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong Credentials."})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong Credentials."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Login Sucessful!",
		"token":   token,
	})

}
