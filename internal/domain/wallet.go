package domain

import "github.com/shopspring/decimal"

type Wallet struct {
	Id      string
	Balance decimal.Decimal
}
