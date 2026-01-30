package main

import (
	"github.com/emmanuel/rest_project/db"
	"github.com/emmanuel/rest_project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") //localhost:8080

}
