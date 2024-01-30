package repository

import (
	"github.com/Warh40k/infotecs_task/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type WalletPostgres struct {
	db *sqlx.DB
}

func (w WalletPostgres) CreateWallet() (domain.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (w WalletPostgres) GetWallet(id string) (domain.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (w WalletPostgres) ShowHistory(id string) ([]domain.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (w WalletPostgres) SendMoney(from, to string, amount decimal.Decimal) error {
	//TODO implement me
	panic("implement me")
}

func NewWalletPostgres(db *sqlx.DB) *WalletPostgres {
	return &WalletPostgres{db: db}
}
