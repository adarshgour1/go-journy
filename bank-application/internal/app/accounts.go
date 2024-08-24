package app

import (
	"context"
)

const (
	createAccountQuery  = `INSERT INTO accounts (owner, balance, currency ) VALUES (?, ?, ?)`
	selectAccountQuery  = `SELECT id, owner, balance, currency, created_at FROM accounts WHERE id = ?`
	selectAccountsQuery = `SELECT id, owner, balance, currency, createTime FROM accounts`
	updateAccountQuery  = `UPDATE accounts SET id = ?,owner = ?, balance = ?, currency = ?, WHERE id = ?`
	deleteAccountQuery  = `DELETE FROM accounts WHERE id = ?`
)

type CreateAccountParams struct {
	Owner    string
	Balance  int
	Currency string
}

func (a *App) CreateAccount(ctx context.Context, createAccountParams CreateAccountParams) (id int64, err error) {
	result, err := a.db.ExecContext(ctx,
		createAccountQuery,
		createAccountParams.Owner,
		createAccountParams.Balance,
		createAccountParams.Currency,
	)
	if err != nil {
		a.log.Printf("error occured while creating account: %s", err.Error())
		return
	}

	id, err = result.LastInsertId()
	if err != nil {
		a.log.Printf("error occured while getting id for account: %s", err.Error())
		return
	}

	return

}

func (a *App) GetAccount(ctx context.Context, accountId int64) (account Account, err error) {
	row := a.db.QueryRowContext(ctx, selectAccountQuery, accountId)

	if err = row.Scan(
		&account.Id,
		&account.Owner,
		&account.Balance,
		&account.Currency,
		&account.CreateTime,
	); err != nil {
		a.log.Printf("error occured while getting account for id %d: %s", accountId, err.Error())
		return
	}

	return
}

func (a *App) GetAccounts(ctx context.Context) (accounts []Account, err error) {
	row, err := a.db.QueryContext(ctx, selectAccountsQuery)
	if err != nil {
		a.log.Printf("error occured while getting accounts: %s", err.Error())
	}

	for row.Next() {
		var account Account
		if err = row.Scan(
			&account.Id,
			&account.Owner,
			&account.Balance,
			&account.Currency,
			&account.CreateTime,
		); err != nil {
			a.log.Printf("error occured while getting account for id %d: %s", account.Id, err.Error())
			return
		}
		accounts = append(accounts, account)
	}

	return
}

type UpdateAccountParams struct {
	Id       int
	Owner    string
	Balance  int
	Currency string
}

func (a *App) UpdateAccount(ctx context.Context, accountId int, account UpdateAccountParams) (id int, err error) {
	if _, err = a.db.ExecContext(
		ctx,
		updateAccountQuery,
		account.Id,
		account.Owner,
		account.Balance,
		account.Currency,
		accountId,
	); err != nil {
		a.log.Printf("error occured while updating accounts: %s", err.Error())
		return
	}

	id = accountId
	return
}

func (a *App) DeleteAccount(ctx context.Context, accountId int) (id int, err error) {
	if _, err = a.db.ExecContext(
		ctx,
		deleteAccountQuery,
		accountId,
	); err != nil {
		a.log.Printf("error occured while deleting accounts: %s", err.Error())
		return
	}

	id = accountId
	return
}
