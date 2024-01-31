package repository

import (
	"fmt"
	"github.com/Warh40k/infotecs_task/internal/domain"
	"github.com/jmoiron/sqlx"
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

	if err != nil {
		return wallet, &domain.NotFoundError{Message: "No wallet with the specified id"}
	}

	return wallet, nil
}

func (r WalletRepository) ShowHistory(id string) ([]domain.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (r WalletRepository) SendMoney(tr domain.Transaction) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	fromWallet, err := getWalletTx(tx, tr.From)
	if err != nil {
		return err
	}

	toWallet, err := getWalletTx(tx, tr.To)
	if err != nil {
		return err
	}

	fromWallet.Balance = fromWallet.Balance.Sub(tr.Amount)
	toWallet.Balance = toWallet.Balance.Add(tr.Amount)

	if err = updateBalanceTx(tx, fromWallet); err != nil {
		tx.Rollback()
		return err
	}

	if err = updateBalanceTx(tx, toWallet); err != nil {
		tx.Rollback()
		return err
	}

	addTransactionQuery := fmt.Sprintf(`INSERT INTO %s("from","to",amount) VALUES($1,$2,$3)`, transactionsTable)
	if _, err = tx.Exec(addTransactionQuery, tr.From, tr.To, tr.Amount); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func getWalletTx(tx *sqlx.Tx, id string) (domain.Wallet, error) {
	var wallet domain.Wallet
	getWalletQuery := fmt.Sprintf(`SELECT * FROM %s WHERE id=$1`, walletsTable)
	err := tx.Get(&wallet, getWalletQuery, id)
	return wallet, err
}

func updateBalanceTx(tx *sqlx.Tx, wallet domain.Wallet) error {
	query := fmt.Sprintf("UPDATE %s SET balance=$1 WHERE id=$2", walletsTable)
	_, err := tx.Exec(query, wallet.Balance, wallet.Id)
	return err
}

func NewWalletRepository(db *sqlx.DB) *WalletRepository {
	return &WalletRepository{db: db}
}
