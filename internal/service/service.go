package service

import (
	"github.com/Warh40k/infotecs_task/internal/repository"
)

type Wallet interface {
}

type Service struct {
	Wallet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Wallet: nil,
	}
}
