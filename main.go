package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin engine
	r := gin.Default()

	// Serve static files from the "static" directory
	r.Static("/static", "./static")

	// Define a route handler for the root URL "/"
	r.GET("/", func(c *gin.Context) {
		// Render the HTML file index.html
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Run the server
	r.Run(":8080")
}
