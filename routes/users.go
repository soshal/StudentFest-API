package routes

import (
	"college/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func  signup(c *gin.Context){

	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data",
		})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}






}


func  login(c *gin.Context){
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data",
		})
		return
	} 





}
