package main

import (
	"net/http"
	"strconv"

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
	server.GET("/events/:id", getEventById)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
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

func getEventById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	event := models.GetEventById(id)
	if event.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	event := models.GetEventById(id)
	if event.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedEvent := models.UpdateEvent(id, event)
	context.JSON(http.StatusOK, updatedEvent)
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	models.DeleteEvent(id)
	context.JSON(http.StatusOK, gin.H{"message": "event deleted"})
}
