package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucianoZgabriel/event-booking/db"
	"github.com/lucianoZgabriel/event-booking/models"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvent)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvent(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse requested data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
