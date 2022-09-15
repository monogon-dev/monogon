-- name: GetAccountByOIDC :many
SELECT
    accounts.*
FROM accounts
WHERE account_oidc_sub = $1;

-- name: InitializeAccountFromOIDC :one
INSERT INTO accounts (
    account_oidc_sub, account_display_name
) VALUES (
    $1, $2
)
RETURNING *;