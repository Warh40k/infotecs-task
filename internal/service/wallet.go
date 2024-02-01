package service

import (
	"errors"
	"fmt"
	"github.com/Warh40k/infotecs_task/internal/app"
	"github.com/Warh40k/infotecs_task/internal/domain"
	"github.com/Warh40k/infotecs_task/internal/repository"
	"github.com/sirupsen/logrus"
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

func (s WalletService) GetWalletHistory(walletId string) ([]domain.Transaction, error) {
	return s.repo.GetWalletHistory(walletId)
}

func (s WalletService) SendMoney(tr domain.Transaction) error {
	if err := s.repo.SendMoney(tr); err != nil {
		logrus.WithError(fmt.Errorf("error sending money: %w",
			errors.Unwrap(err))).Errorf("Error to process transaction %+v", tr)

		var badRequest app.BadRequestError
		var notFound app.NotFoundError
		if errors.As(err, &badRequest) || errors.As(err, &notFound) {
			return err
		} else {
			return app.InternalError{Message: "Unhandled error", Err: err}
		}
	}

	return nil
}

func NewWalletService(repo repository.Wallet) *WalletService {
	return &WalletService{repo: repo}
}
