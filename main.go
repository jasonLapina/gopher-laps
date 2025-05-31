package main

import (
	"awesomeProject/database"
	"awesomeProject/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDb()
	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
