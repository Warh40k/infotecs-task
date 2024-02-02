package repository

import (
	"fmt"
	"github.com/Warh40k/infotecs_task/internal/app"
	"github.com/Warh40k/infotecs_task/internal/domain"
	"github.com/jmoiron/sqlx"
)

type WalletRepository struct {
	db *sqlx.DB
}

// Create default wallet
func (r WalletRepository) CreateWallet() (domain.Wallet, error) {
	var wallet domain.Wallet
	id, err := GenerateId()
	if err != nil {
		return wallet, app.InternalError{Message: "error generating id"}
	}

	query := fmt.Sprintf("INSERT INTO %s(id,balance) VALUES($1,$2) RETURNING id, balance", walletsTable)
	err = r.db.Get(&wallet, query, id, defaultBalance)

	return wallet, err
}

// Get wallet data
func (r WalletRepository) GetWallet(id string) (domain.Wallet, error) {
	var wallet domain.Wallet

	query := fmt.Sprintf("SELECT * from %s WHERE id=$1", walletsTable)
	err := r.db.Get(&wallet, query, id)

	if err != nil {
		return wallet, &app.NotFoundError{Message: "no wallet with the specified id"}
	}

	return wallet, nil
}

func (r WalletRepository) GetWalletHistory(walletId string) ([]domain.Transaction, error) {
	var trs []domain.Transaction
	if _, err := r.GetWallet(walletId); err != nil {
		return trs, &app.NotFoundError{Message: "no wallet with the specified id"}
	}
	query := fmt.Sprintf(`SELECT * FROM %s WHERE "from"=$1 OR "to"=$1`, transactionsTable)
	err := r.db.Select(&trs, query, walletId)

	return trs, err
}

// Process given transaction
func (r WalletRepository) SendMoney(tr domain.Transaction) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	fromWallet, err := getWalletTx(tx, tr.From)
	if err != nil {
		return app.BadRequestError{Message: "error getting sender's wallet", Err: err}
	}

	toWallet, err := getWalletTx(tx, tr.To)
	if err != nil {
		return app.NotFoundError{Message: "error getting receiver's wallet", Err: err}
	}

	fromWallet.Balance -= tr.Amount
	toWallet.Balance += tr.Amount

	if err = updateBalanceTx(tx, fromWallet); err != nil {
		tx.Rollback()
		return app.BadRequestError{Message: "error updating sender's balance", Err: err}
	}

	if err = updateBalanceTx(tx, toWallet); err != nil {
		tx.Rollback()
		return app.InternalError{Message: "error updating receiver's balance", Err: err}
	}

	addTransactionQuery := fmt.Sprintf(`INSERT INTO %s("id","from","to","amount") VALUES($1,$2,$3,$4)`, transactionsTable)
	id, err := GenerateId()
	if err != nil {
		return app.InternalError{Message: "error generating id"}
	}
	if _, err = tx.Exec(addTransactionQuery, id, tr.From, tr.To, tr.Amount); err != nil {
		tx.Rollback()
		return app.InternalError{Message: "error saving transaction", Err: err}
	}

	return tx.Commit()
}

// Get wallet during transaction
func getWalletTx(tx *sqlx.Tx, id string) (domain.Wallet, error) {
	var wallet domain.Wallet
	getWalletQuery := fmt.Sprintf(`SELECT * FROM %s WHERE id=$1`, walletsTable)
	err := tx.Get(&wallet, getWalletQuery, id)
	return wallet, err
}

// Update balance during transaction
func updateBalanceTx(tx *sqlx.Tx, wallet domain.Wallet) error {
	query := fmt.Sprintf("UPDATE %s SET balance=$1 WHERE id=$2", walletsTable)
	_, err := tx.Exec(query, wallet.Balance, wallet.Id)
	return err
}

func NewWalletRepository(db *sqlx.DB) *WalletRepository {
	return &WalletRepository{db: db}
}
