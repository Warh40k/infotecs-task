package domain

type Wallet struct {
	Id      string  `json:"id" db:"id"`
	Balance float32 `json:"balance" db:"balance"`
}
