package db

import (
	"context"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount (t *testing.T) {

	arg := CreateAccountParams {
		Owner: "admin",
		Balance: 100,
		Currency: "AUD",
	}

	acc, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, acc)
	require.Equal(t, arg.Owner, acc.Owner)
	require.Equal(t, arg.Balance, acc.Balance)
	require.Equal(t, arg.Currency, acc.Currency)

	require.NotZero(t,acc.ID)
	require.NotZero(t, acc.CreatedAt)
}

func TestGetAccount (t *testing.T) {
	acc, err := testQueries.GetAccount(context.Background(), 1)

	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, acc.Balance, int64(100))

}