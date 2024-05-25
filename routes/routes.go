package routes


import (
	


	"github.com/gin-gonic/gin"
)


func ResigterEventRoutes(r *gin.Engine) {
	r.GET("/events", eventManager)
	r.GET("/events/:id", eventManager)
	r.POST("/events", contentManager)
	r.PUT("/events/:id", updateManager)
	r.DELETE("/events/:id", deleteManager)
	r.POST("/login", signup)
}