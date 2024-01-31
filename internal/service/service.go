package service

import (
	"github.com/Warh40k/infotecs_task/internal/domain"
	"github.com/Warh40k/infotecs_task/internal/repository"
	"github.com/shopspring/decimal"
)

type Wallet interface {
	CreateWallet() (*domain.Wallet, error)
	GetWallet(id string) (*domain.Wallet, error)
	ShowHistory(id string) ([]domain.Transaction, error)
	SendMoney(from, to string, amount decimal.Decimal) error
}

type Service struct {
	Wallet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Wallet: NewWalletService(repos.Wallet),
	}
}
