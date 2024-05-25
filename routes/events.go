package routes


import (
	                  // Import the db package
	"college/models" // Import the models package
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)






func eventManager(c *gin.Context) {
	
	events , err := models.GetEvents() // Use models.GetEvents instead of event.GetEvents

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"events": events,
	})
   
}

func eventManagerbyId(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}
      _, err = models.GetEventbyId(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Event not found",
		})
		return
	}

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



func updateManager(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}

	_, err = models.GetEventbyId(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Event not found",
		})
		return
	}

	var update models.Event
	err = c.ShouldBindJSON(&update)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data",
		})
		return
	}

	update.Id = id
	
	
	update.UpdateEvent()
	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated",
		"event":   update,
	})





}


func deleteManager(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}

	_, err = models.GetEventbyId(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Event not found",
		})
		return
	}

	var delete models.Event
	err = c.ShouldBindJSON(&delete)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data",
		})
		return
	}

	delete.Id = id
	delete.DeleteEvent()
	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted",
		"event":   delete,
	})
	

}




	


	