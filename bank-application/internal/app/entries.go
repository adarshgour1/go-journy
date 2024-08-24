package app

import (
	"context"
)

const (
	createEntryQuery   = `INSERT INTO entries ( account_id, amount ) VALUES ( ?, ? )`
	selectEntryQuery   = `SELECT id, account_id, ammount, created_at FROM entries WHERE id = ?`
	selectEntriesQuery = `SELECT id, account_id, ammount, created_at FROM entries`
	updateEntryQuery   = `UPDATE entries SET id = ?, account_id = ?, ammount = ?, WHERE id = ?`
	deleteEntryQuery   = `DELETE FROM entries WHERE id = ?`
)

type CreateEntryParams struct {
	AccountId int
	Amount    int
}

func (a *App) CreateEntry(ctx context.Context, createEntryParams CreateEntryParams) (id int64, err error) {
	result, err := a.db.ExecContext(ctx,
		createEntryQuery,
		createEntryParams.AccountId,
		createEntryParams.Amount,
	)
	if err != nil {
		a.log.Printf("error occured while creating entry: %s", err.Error())
		return
	}

	id, err = result.LastInsertId()
	if err != nil {
		a.log.Printf("error occured while getting id for entry: %s", err.Error())
		return
	}

	return

}

func (a *App) GetEntry(ctx context.Context, entryId int) (entry Entry, err error) {
	row := a.db.QueryRowContext(ctx, selectEntryQuery, entryId)

	if err = row.Scan(
		&entry.Id,
		&entry.AccountId,
		&entry.Amount,
		&entry.CreateTime,
	); err != nil {
		a.log.Printf("error occured while getting entry for id %d: %s", entryId, err.Error())
		return
	}

	return
}

func (a *App) GetEntries(ctx context.Context) (entries []Entry, err error) {
	row, err := a.db.QueryContext(ctx, selectEntriesQuery)
	if err != nil {
		a.log.Printf("error occured while getting entries: %s", err.Error())
	}

	for row.Next() {
		var entry Entry
		if err = row.Scan(
			&entry.Id,
			&entry.AccountId,
			&entry.Amount,
			&entry.CreateTime,
		); err != nil {
			a.log.Printf("error occured while getting entry for id %d: %s", entry.Id, err.Error())
			return
		}
		entries = append(entries, entry)
	}

	return
}

type UpdateEntryParams struct {
	Id        int
	AccountId int
	Amount    int
}

func (a *App) UpdateEntry(ctx context.Context, entryId int, entry UpdateEntryParams) (id int, err error) {
	if _, err = a.db.ExecContext(
		ctx,
		updateEntryQuery,
		entry.Id,
		entry.AccountId,
		entry.Amount,
		entryId,
	); err != nil {
		a.log.Printf("error occured while updating entry: %s", err.Error())
		return
	}

	id = entryId
	return
}

func (a *App) DeleteEntry(ctx context.Context, entryId int) (id int, err error) {
	if _, err = a.db.ExecContext(
		ctx,
		deleteEntryQuery,
		entryId,
	); err != nil {
		a.log.Printf("error occured while deleting entry for id %d: %s", entryId, err.Error())
		return
	}

	id = entryId
	return
}
