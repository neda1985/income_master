package handlers

import (
	"github.com/gin-gonic/gin"
	"incomeCoach/Transaction/models"

	"net/http"
)

func AddMoneyHandler(c *gin.Context) {
	var input models.InsertMoney

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Money added successfully"})
}

func SpendMoneyHandler(c *gin.Context) {
	var input models.GetMoney

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Money decreased successfully"})
}

func LeftMoneyHandler(c *gin.Context) {
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
