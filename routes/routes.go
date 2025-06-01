package routes

import (
	"awesomeProject/midwares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	protected := server.Group("/")
	protected.Use(midwares.Authenticate)
	protected.POST("/events", createEvent)
	protected.PUT("/events/:id", updateEvent)
	protected.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
