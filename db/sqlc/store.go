package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	FromID int64 `json:"from_id"`
	ToID   int64 `json:"to_id"`
	Amount int64 `json:"amount"`
}

type TransferTxResults struct {
	Transfer  Transfer `json:"transfer"`
	FromEntry Entry    `json:"from_ent"`
	ToEntry   Entry    `json:"to_ent"`
}

func (store *Store) TransferTx(ctx context.Context, args TransferTxParams) (TransferTxResults, error) {
	var res TransferTxResults

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		res.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: args.FromID,
			ToAccountID:   args.ToID,
			Amount:        args.Amount,
		})

		if err != nil {
			return err
		}

		res.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			Amount:    -args.Amount,
			AccountID: args.FromID,
		})

		if err != nil {
			return err
		}

		res.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: args.ToID,
			Amount:    args.Amount,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return res, err
}
