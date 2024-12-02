package persistence

import (
	"context"
	db "github.com/jpmoraess/backend-challenge-go/db/sqlc"
	"github.com/jpmoraess/backend-challenge-go/internal/domain"
)

type walletRepositoryAdapter struct {
	store db.Store
}

func NewWalletRepositoryAdapter(store db.Store) *walletRepositoryAdapter {
	return &walletRepositoryAdapter{store: store}
}

func (r *walletRepositoryAdapter) Get(ctx context.Context, id int64) (wallet *domain.Wallet, err error) {
	return nil, err
}

func (r *walletRepositoryAdapter) Save(ctx context.Context, wallet *domain.Wallet) (err error) {
	_, err = r.store.CreateWallet(ctx, db.CreateWalletParams{
		Type:     wallet.WalletType(),
		FullName: wallet.FullName(),
		Document: wallet.Document(),
		Email:    wallet.Email(),
		Password: wallet.Password(),
		Balance:  wallet.Balance(),
	})
	if err != nil {
		return
	}

	return
}
