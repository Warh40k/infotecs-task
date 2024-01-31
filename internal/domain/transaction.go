package domain

import (
	"github.com/shopspring/decimal"
	"time"
)

type Transaction struct {
	Id     string          `json:"id" db:"id"`
	From   string          `json:"from" db:"from"`
	To     string          `json:"to" db:"to"`
	Amount decimal.Decimal `json:"amount" db:"amount"`
	Time   time.Time       `json:"time" db:"time"`
}
