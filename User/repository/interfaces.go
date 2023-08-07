package repository

import (
	"context"
	"time"
)

type MoneyRepository interface {
	AddMoney(ctx context.Context, userId string, amount int, createdAt time.Time) (int, error)
	SpendMoney(tx context.Context, userId string, amount int, createdAt time.Time) (int, error)
	restMoney(ctx context.Context, userId string) (int error)
}
