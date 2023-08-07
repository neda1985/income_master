package models

import "time"

type Transaction struct {
	Amount    int           `bson:"amount"`
	Category  string        `bson:"category"`
	CreatedAt time.Time     `bson:"created_at"`
	Location  time.Location `bson:"location"`
}
