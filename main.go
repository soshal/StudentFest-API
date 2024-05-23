package main

import (
	"net/http"
	"college/db" // Import the db package
	"github.com/gin-gonic/gin"
	"college/models" // Import the models package
)

func main() {
	db.InitDB() // Initialize the database
	r := gin.Default()

	r.GET("/events", eventManager)
	r.POST("/events", contentManager)
	r.Run(":5500")
}

func eventManager(c *gin.Context) {
	
	events := models.GetEvents() // Use models.GetEvents instead of event.GetEvents
	c.JSON(http.StatusOK, gin.H{
		"events": events,
	})
   
}

func contentManager(c *gin.Context) {
	 var event models.Event
	 err := c.ShouldBindJSON(&event)
	
	 if err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{
			 "error": "Invalid data",
		 })
		}
	
	 event.Save()

	 c.JSON(http.StatusCreated, gin.H{
		 "message": "Event created", "event": event,

	 })
	
}
