package app

import (
	"context"
)

const (
	createTransferQuery  = `INSERT INTO transfers ( from_account_id, to_account_id, amount ) VALUES ( ?, ?, ? )`
	selectTransferQuery  = `SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers WHERE id = ?`
	selectTransfersQuery = `SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers`
	updateTransferQuery  = `UPDATE transfers SET id = ?, from_account_id = ?, to_account_id = ?, amount = ? WHERE id = ?`
	deleteTransferQuery  = `DELETE FROM transfers WHERE id = ?`
)

type CreateTransferParams struct {
	FromAccountId string
	ToAccountId   int
	Amount        string
}

func (a *App) CreateTransfer(ctx context.Context, createTransferParams CreateTransferParams) (id int64, err error) {
	result, err := a.db.ExecContext(ctx,
		createTransferQuery,
		createTransferParams.FromAccountId,
		createTransferParams.ToAccountId,
		createTransferParams.Amount,
	)
	if err != nil {
		a.log.Printf("error occured while creating transfer: %s", err.Error())
		return
	}

	id, err = result.LastInsertId()
	if err != nil {
		a.log.Printf("error occured while getting id for transfer: %s", err.Error())
		return
	}

	return

}

func (a *App) GetTransfer(ctx context.Context, transferId int64) (transfer Transfer, err error) {
	row := a.db.QueryRowContext(ctx, selectTransferQuery, transferId)

	if err = row.Scan(
		&transfer.Id,
		&transfer.FromAccountId,
		&transfer.ToAccountId,
		&transfer.Amount,
		&transfer.CreateTime,
	); err != nil {
		a.log.Printf("error occured while getting transer for id %d: %s", transferId, err.Error())
		return
	}

	return
}

func (a *App) GetTransfers(ctx context.Context) (transfers []Transfer, err error) {
	row, err := a.db.QueryContext(ctx, selectTransfersQuery)
	if err != nil {
		a.log.Printf("error occured while getting transfers: %s", err.Error())
	}

	for row.Next() {
		var transfer Transfer
		if err = row.Scan(
			&transfer.Id,
			&transfer.FromAccountId,
			&transfer.ToAccountId,
			&transfer.Amount,
			&transfer.CreateTime,
		); err != nil {
			a.log.Printf("error occured while getting transfer for id %d: %s", transfer.Id, err.Error())
			return
		}
		transfers = append(transfers, transfer)
	}

	return
}

type UpdateTransferParams struct {
	Id            int
	FromAccountId int
	ToAccountId   int
	Amount        int
}

func (a *App) UpdateTransfer(ctx context.Context, transferId int, transfer UpdateTransferParams) (id int, err error) {
	if _, err = a.db.ExecContext(
		ctx,
		updateTransferQuery,
		transfer.Id,
		transfer.FromAccountId,
		transfer.ToAccountId,
		transfer.Amount,
		transferId,
	); err != nil {
		a.log.Printf("error occured while updating transfer: %s", err.Error())
		return
	}

	id = transferId
	return
}

func (a *App) DeleteTransfer(ctx context.Context, transferId int) (id int, err error) {
	if _, err = a.db.ExecContext(
		ctx,
		deleteTransferQuery,
		transferId,
	); err != nil {
		a.log.Printf("error occured while deleting transfer: %s", err.Error())
		return
	}

	id = transferId
	return
}
