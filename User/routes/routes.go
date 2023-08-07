package routes

import (
	"github.com/gin-gonic/gin"
	"incomeCoach/handlers"
)

func SetUp(router *gin.Engine)  {
	router.POST("/add", handlers.AddMoneyHandler)
	router.POST("/spend",handlers.SpendMoneyHandler)
	router.GET("/rest",handlers.LeftMoneyHandler)
}
