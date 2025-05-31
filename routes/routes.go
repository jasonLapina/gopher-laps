package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	//GET EVENTS
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)

	//UPDATE EVENTS
	server.PUT("/events/:id")

	//POST EVENTS
	server.POST("/events", CreateEvent)
}
