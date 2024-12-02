-- name: CreateWallet :one
INSERT INTO wallets (type, full_name, document, email, password)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetWallet :one
SELECT * FROM wallets
WHERE id = $1 LIMIT 1;

-- name: GetWalletForUpdate :one
SELECT * FROM wallets
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateWallet :one
UPDATE wallets
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: AddBalanceToWallet :one
UPDATE wallets
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;