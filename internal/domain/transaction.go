package domain

import (
	"encoding/json"
	"time"
)

type Transaction struct {
	Time   time.Time `json:"time" db:"time"`
	Id     string    `json:"-" db:"id"`
	From   string    `json:"from" db:"from"`
	To     string    `json:"to" db:"to"`
	Amount float32   `json:"amount" db:"amount" binding:"gte=0"`
}

func (t Transaction) MarshalJSON() ([]byte, error) {
	formattedTime := t.Time.Format(time.RFC3339)
	return json.Marshal(struct {
		Time   string  `json:"time"`
		From   string  `json:"from"`
		To     string  `json:"to"`
		Amount float32 `json:"amount"`
	}{
		Time:   formattedTime,
		From:   t.From,
		To:     t.To,
		Amount: t.Amount,
	})
}
