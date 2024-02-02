package domain

import (
	"time"
)

type Transaction struct {
	Time   time.Time `json:"time" db:"time"`
	Id     string    `json:"-" db:"id"`
	From   string    `json:"from" db:"from"`
	To     string    `json:"to" db:"to"`
	Amount float32   `json:"amount" db:"amount" binding:"gte=0"`
}
