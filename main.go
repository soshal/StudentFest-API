package main

import (
	"college/db"     // Import the db package
	 // Import the models package
	"college/routes"
	

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() // Initialize the database
	r := gin.Default()

	routes.ResigterEventRoutes(r)
	
	r.Run(":5500")
}

