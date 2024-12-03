package domain

import (
	"context"
	"errors"
)

type Wallet struct {
	id         int64
	walletType string
	fullName   string
	document   string
	email      string
	password   string
	balance    int64
}

type WalletRepository interface {
	Get(ctx context.Context, id int64) (wallet *Wallet, err error)
	Save(ctx context.Context, wallet *Wallet) (err error)
}

func NewWallet(walletType string, fullName string, document string, email string, password string) (w *Wallet, err error) {
	w = &Wallet{
		walletType: walletType,
		fullName:   fullName,
		document:   document,
		email:      email,
		password:   password,
		balance:    0,
	}
	if err = w.validate(); err != nil {
		return
	}

	return
}

func RestoreWallet(id int64, walletType string, fullName string, document string, email string, password string, balance int64) (w *Wallet, err error) {
	w = &Wallet{
		id:         id,
		walletType: walletType,
		fullName:   fullName,
		document:   document,
		email:      email,
		password:   password,
		balance:    balance,
	}
	if err = w.validate(); err != nil {
		return
	}

	return
}

func (w *Wallet) validate() error {
	if len(w.fullName) == 0 {
		return errors.New("full name is required")
	}
	if len(w.document) == 0 {
		return errors.New("document is required")
	}
	if len(w.email) == 0 {
		return errors.New("email is required")
	}
	if len(w.password) == 0 {
		return errors.New("password is required")
	}
	return nil
}

func (w *Wallet) ID() int64 {
	return w.id
}

func (w *Wallet) WalletType() string {
	return w.walletType
}

func (w *Wallet) FullName() string {
	return w.fullName
}

func (w *Wallet) Document() string {
	return w.document
}

func (w *Wallet) Email() string {
	return w.email
}

func (w *Wallet) Password() string {
	return w.password
}

func (w *Wallet) Balance() int64 {
	return w.balance
}
