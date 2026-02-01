package routes

import (
	"github.com/emmanuel/rest_project/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/event/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/event/:id", updateEvent)
	authenticated.DELETE("/event/:id", deleteEvent)
	authenticated.POST("/event/:id/register", RegisterforEvents)
	authenticated.DELETE("/event/:id/register", CancelRegistrationForEvents)

	// ============Users API=========
	server.POST("/signup", signup)
	server.POST("/login", login)

}
