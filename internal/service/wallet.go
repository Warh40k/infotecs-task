package service

import (
	"github.com/Warh40k/infotecs_task/internal/domain"
	"github.com/Warh40k/infotecs_task/internal/repository"
	"github.com/shopspring/decimal"
)

type WalletService struct {
	repo repository.Wallet
}

func (s WalletService) CreateWallet() (domain.Wallet, error) {
	return s.repo.CreateWallet()
}

func (s WalletService) GetWallet(id string) (domain.Wallet, error) {
	return s.repo.GetWallet(id)
}

func (s WalletService) ShowHistory(id string) ([]domain.Transaction, error) {
	return s.repo.ShowHistory(id)
}

func (s WalletService) SendMoney(from, to string, amount decimal.Decimal) error {
	return s.repo.SendMoney(from, to, amount)
}

func NewWalletService(repo repository.Wallet) *WalletService {
	return &WalletService{repo: repo}
}
