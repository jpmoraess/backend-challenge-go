package db

import (
	"context"
)

// TransferTxParam contains the input params of the transfer tx
type TransferTxParam struct {
	FromWalletID int64 `json:"from_wallet_id"`
	ToWalletID   int64 `json:"to_wallet_id"`
	Amount       int64 `json:"amount"`
}

// TransferTxResult contains the result of transfer tx
type TransferTxResult struct {
	Transfer   Transfer `json:"transfer"`
	FromWallet Wallet   `json:"from_wallet"`
	ToWallet   Wallet   `json:"to_wallet"`
	FromEntry  Entry    `json:"from_entry"`
	ToEntry    Entry    `json:"to_entry"`
}

func (s *SQLStore) TransferTx(ctx context.Context, arg TransferTxParam) (result TransferTxResult, err error) {
	err = s.execTx(ctx, func(q *Queries) error {
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromWallet: arg.FromWalletID,
			ToWallet:   arg.ToWalletID,
			Amount:     arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			WalletID: arg.FromWalletID,
			Amount:   -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			WalletID: arg.ToWalletID,
			Amount:   arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.FromWalletID < arg.ToWalletID {
			result.FromWallet, result.ToWallet, err =
				addMoney(ctx, q, arg.FromWalletID, -arg.Amount, arg.ToWalletID, arg.Amount)
		} else {
			result.ToWallet, result.FromWallet, err =
				addMoney(ctx, q, arg.ToWalletID, arg.Amount, arg.FromWalletID, -arg.Amount)
		}

		return err
	})

	return result, err
}

func addMoney(
	ctx context.Context,
	q *Queries,
	walletID1,
	amount1,
	walletID2,
	amount2 int64,
) (wallet1 Wallet, wallet2 Wallet, err error) {
	wallet1, err = q.AddBalanceToWallet(ctx, AddBalanceToWalletParams{
		ID:     walletID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	wallet2, err = q.AddBalanceToWallet(ctx, AddBalanceToWalletParams{
		ID:     walletID2,
		Amount: amount2,
	})

	return
}
