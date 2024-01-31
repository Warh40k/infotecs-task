package repository

import (
	"github.com/Warh40k/infotecs_task/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Wallet interface {
	CreateWallet() (domain.Wallet, error)
	GetWallet(id string) (domain.Wallet, error)
	ShowHistory(id string) ([]domain.Transaction, error)
	SendMoney(tr domain.Transaction) error
}

type Repository struct {
	Wallet
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Wallet: NewWalletRepository(db),
	}
}
