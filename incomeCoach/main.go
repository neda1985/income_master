package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router instance
	router := gin.Default()

	// Define your endpoints and handlers
	router.GET("/endpoint1", endpoint1Handler)
	router.POST("/endpoint2", endpoint2Handler)

	// Start the server
	router.Run(":8080")
}
func endpoint1Handler(c *gin.Context) {
	// Get parameters from the query string
	param1 := c.Query("param1")
	param2 := c.Query("param2")

	// You can perform any logic or processing here
	// For now, let's just return the received parameters as a response
	c.JSON(200, gin.H{
		"message": "Endpoint 1",
		"param1":  param1,
		"param2":  param2,
	})
}

func endpoint2Handler(c *gin.Context) {
	// Create a struct to bind the parameters from the request body
	type RequestParams struct {
		Param3 string `json:"param3"`
		Param4 int    `json:"param4"`
	}

	var requestParams RequestParams
	err := c.ShouldBindJSON(&requestParams)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON data"})
		return
	}

	// You can perform any logic or processing here
	// For now, let's just return the received parameters as a response
	c.JSON(200, gin.H{
		"message": "Endpoint 2",
		"param3":  requestParams.Param3,
		"param4":  requestParams.Param4,
	})
}
