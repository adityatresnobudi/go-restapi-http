package transaction_pg

const GET_ALL_TRANSACTIONS = `
	SELECT id, account_number, transaction_type, amount, created_at, updated_at
	from transactions
`

const GET_ONE_TRANSACTION_BY_ID = `
	SELECT id, account_number, transaction_type, amount, created_at, updated_at
	from transactions WHERE id = $1
`

const GET_ONE_TRANSACTION_BY_ACCOUNT_NUMBER = `
	SELECT id, account_number, transaction_type, amount, created_at, updated_at
	from transactions WHERE account_number = $1
`

const INSERT_TRANSACTION = `
	INSERT INTO transactions (account_number, transaction_type, amount) 
	VALUES ($1, $2, $3)
	RETURNING id, account_number, transaction_type, amount, created_at, updated_at
`

const UPDATE_TRANSACTION = `
	UPDATE transactions
	SET account_number = $1, transaction_type = $2, amount = $3
	WHERE id = $4
	RETURNING id, account_number, transaction_type, amount, updated_at
`

const DELETE_TRANSACTION = `
	DELETE FROM transactions
	WHERE id = $1
`
