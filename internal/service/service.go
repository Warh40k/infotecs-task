package service

import (
	"github.com/Warh40k/infotecs_task/internal/domain"
	"github.com/Warh40k/infotecs_task/internal/repository"
)

type Wallet interface {
	CreateWallet() (domain.Wallet, error)
	GetWallet(id string) (domain.Wallet, error)
	GetWalletHistory(walletId string) ([]domain.Transaction, error)
	SendMoney(tr domain.Transaction) error
}

type Service struct {
	Wallet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Wallet: NewWalletService(repos.Wallet),
	}
}
