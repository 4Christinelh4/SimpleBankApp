package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	acc1 := CreateRandomAccount(t)

	ent, err := testQueries.CreateEntry(context.Background(), CreateEntryParams{
		AccountID: acc1.ID,
		Amount: acc1.Balance,
	})

	require.NoError(t, err)
	require.NotEmpty(t, ent)
	require.Equal(t, ent.Amount, acc1.Balance)
}