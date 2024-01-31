package domain

import "github.com/shopspring/decimal"

type Wallet struct {
	Id      string          `json:"id" db:"id"`
	Balance decimal.Decimal `json:"balance" db:"balance"`
}
