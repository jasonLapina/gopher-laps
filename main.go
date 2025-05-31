package main

import (
	"awesomeProject/database"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	database.InitDb()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func getEvents(context *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}
	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusCreated, event)
}
