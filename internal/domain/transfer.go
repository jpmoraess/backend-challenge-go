package domain

type Transfer struct {
	ID           int64
	FromWalletId int64
	ToWalletId   int64
	Amount       int64
}
