package domain

import "github.com/shopspring/decimal"

type Transaction struct {
	Id     string
	From   Wallet
	To     Wallet
	Amount decimal.Decimal
}
