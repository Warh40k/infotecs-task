package domain

type Wallet struct {
	Id      string  `json:"id,string" db:"id"`
	Balance float32 `json:"balance" db:"balance"`
}
