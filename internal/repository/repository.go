package repository

import (
	"github.com/Warh40k/infotecs_task/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type Wallet interface {
	CreateWallet() (*domain.Wallet, error)
	GetWallet(id string) (*domain.Wallet, error)
	ShowHistory(id string) ([]domain.Transaction, error)
	SendMoney(from, to string, amount decimal.Decimal) error
}

type Repository struct {
	Wallet
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Wallet: NewWalletPostgres(db),
	}
}
