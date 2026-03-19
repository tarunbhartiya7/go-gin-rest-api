package main

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	// r := gin.Default()

	// Define a simple GET endpoint
	// r.GET("/ping", func(c *gin.Context) {
	// 	// Return JSON response
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.Run(":8080")

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	// if err := r.Run(); err != nil {
	// 	log.Fatalf("failed to run server: %v", err)
	// }

	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	// context.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdEvent := event.Save()
	context.JSON(http.StatusCreated, createdEvent)
}
