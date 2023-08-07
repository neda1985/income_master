package models

type InsertMoney struct {
	Amount   int    `json:"amount" binding:"required"`
	Category string `json:"category" binding:"required,gt=0"`
}

type GetMoney struct {
	Amount   int    `json:"amount" binding:"required"`
	Category string `json:"category" binding:"required,lt=0"`
}
