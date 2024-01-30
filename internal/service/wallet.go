package service

import (
	"github.com/Warh40k/infotecs_task/internal/domain"
	"github.com/Warh40k/infotecs_task/internal/repository"
	"github.com/shopspring/decimal"
)

type WalletService struct {
	repo repository.Wallet
}

func (w WalletService) CreateWallet() (domain.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (w WalletService) GetWallet(id string) (domain.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (w WalletService) ShowHistory(id string) ([]domain.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (w WalletService) SendMoney(from, to string, amount decimal.Decimal) error {
	//TODO implement me
	panic("implement me")
}

func NewWalletService(repo repository.Wallet) *WalletService {
	return &WalletService{repo: repo}
}
