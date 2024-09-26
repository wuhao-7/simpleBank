-- name: CreateAccount :one
INSERT INTO accounts (
  owner,
  balance,
  currency
) VALUES (
  $1, $2, $3
)RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 limit 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 limit 1
FOR NO KEY UPDATE;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: AddAccountsBalance :one
UPDATE accounts SET balance = balance + sqlc.arg(amount)
where id= sqlc.arg(id)
RETURNING *;

-- name: UpdateAccounts :one
UPDATE accounts SET balance = $2
where id= $1
RETURNING *;

-- name: DeleteAccounts :exec
DELETE FROM accounts WHERE id = $1;
