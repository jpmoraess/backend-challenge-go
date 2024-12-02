package usecase

import (
	"context"
	"github.com/jpmoraess/backend-challenge-go/internal/domain"
)

type createWallet struct {
	walletRepository domain.WalletRepository
}

type CreateWalletInput struct {
	WalletType string
	FullName   string
	Document   string
	Email      string
	Password   string
}

func NewCreateWallet(walletRepository domain.WalletRepository) *createWallet {
	return &createWallet{walletRepository: walletRepository}
}

func (uc *createWallet) Execute(ctx context.Context, input *CreateWalletInput) (err error) {
	wallet, err := domain.NewWallet(input.WalletType, input.FullName, input.Document, input.Email, input.Password)
	if err != nil {
		return
	}

	err = uc.walletRepository.Save(ctx, wallet)
	if err != nil {
		return
	}

	return
}
