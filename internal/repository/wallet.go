package repository

import (
	"fmt"
	"github.com/Warh40k/infotecs_task/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type WalletRepository struct {
	db *sqlx.DB
}

func (r WalletRepository) CreateWallet() (domain.Wallet, error) {
	var wallet domain.Wallet

	query := fmt.Sprintf("INSERT INTO %s(balance) VALUES(%d) RETURNING id, balance", walletsTable, defaultBalance)
	err := r.db.Get(&wallet, query)

	return wallet, err
}

func (r WalletRepository) GetWallet(id string) (domain.Wallet, error) {
	var wallet domain.Wallet

	query := fmt.Sprintf("SELECT * from %s WHERE id=$1", walletsTable)
	err := r.db.Get(&wallet, query, id)

	return wallet, err
}

func (r WalletRepository) ShowHistory(id string) ([]domain.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (r WalletRepository) SendMoney(from, to string, amount decimal.Decimal) error {
	//TODO implement me
	panic("implement me")
}

func NewWalletRepository(db *sqlx.DB) *WalletRepository {
	return &WalletRepository{db: db}
}
