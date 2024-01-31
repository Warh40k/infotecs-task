package repository

import (
	"fmt"
	"github.com/Warh40k/infotecs_task/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type WalletPostgres struct {
	db *sqlx.DB
}

func (r WalletPostgres) CreateWallet() (*domain.Wallet, error) {
	wallet := new(domain.Wallet)

	query := fmt.Sprintf("INSERT INTO %s(balance) VALUES(%d) RETURNING id, balance", walletsTable, defaultBalance)
	err := r.db.Get(wallet, query)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (r WalletPostgres) GetWallet(id string) (*domain.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (r WalletPostgres) ShowHistory(id string) ([]domain.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (r WalletPostgres) SendMoney(from, to string, amount decimal.Decimal) error {
	//TODO implement me
	panic("implement me")
}

func NewWalletPostgres(db *sqlx.DB) *WalletPostgres {
	return &WalletPostgres{db: db}
}
