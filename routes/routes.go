package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	//GET EVENTS
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	//UPDATE EVENTS
	server.PUT("/events/:id", updateEvent)

	//POST EVENTS
	server.POST("/events", createEvent)

	//	DELETE EVENT
	server.DELETE("/events/:id")
}
