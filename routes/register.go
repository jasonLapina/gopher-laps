package routes

import (
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, _ := strconv.ParseInt(context.Param("id"), 10, 64)

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, err)
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered for event"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, _ := strconv.ParseInt(context.Param("id"), 10, 64)

	err := models.CancelRegistration(userId, eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled registration"})

}
