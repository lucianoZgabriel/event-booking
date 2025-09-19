package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucianoZgabriel/event-booking/db"
	"github.com/lucianoZgabriel/event-booking/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
