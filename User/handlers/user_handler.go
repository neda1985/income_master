package handlers

import (
	"github.com/gin-gonic/gin"
	"incomeCoach/User/models"
	"net/http"
)

func UserRegisterHandler(c *gin.Context) {
	var input models.UserRegister

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Money added successfully"})
}
