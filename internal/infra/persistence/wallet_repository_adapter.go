package persistence

import (
	"context"
	"errors"
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
	result, err := r.store.GetWallet(ctx, id)
	if err != nil {

	}

	wallet, err = domain.RestoreWallet(
		result.ID,
		result.Type,
		result.FullName,
		result.Document,
		result.Email,
		result.Password,
		result.Balance,
	)
	if err != nil {
		return
	}

	return
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
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			return errors.New("wallet with document or email already exists")
		}
		return
	}

	return
}
