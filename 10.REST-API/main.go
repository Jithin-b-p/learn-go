package main

import (
	"net/http"

	"github.com/Jithin-b-p/learn-go/10.REST-API/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/events", getEvents)
	r.POST("/events", createEvent)

	r.Run(":8080")
}

func getEvents(context *gin.Context) {

	events := models.GetAllEvents()

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {

	var event models.Event

	if err := context.ShouldBindJSON(&event); err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"message": "value error"})
	} else {

		event.Save()
		context.JSON(http.StatusCreated, gin.H{"message": "Event created"})
	}
}
