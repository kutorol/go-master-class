package db

import (
	"backend-master-class/util"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Accounts {
	arg := CreateAuthorParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCarrency(),
	}

	account, err := testQueries.CreateAuthor(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	a := createRandomAccount(t)
	a2, err := testQueries.GetAccount(context.Background(), a.ID)
	require.NoError(t, err)
	require.NotEmpty(t, a2)

	require.Equal(t, a.ID, a2.ID)
	require.Equal(t, a.Owner, a2.Owner)
	require.Equal(t, a.Balance, a2.Balance)
	require.Equal(t, a.Currency, a2.Currency)
	require.WithinDuration(t, a.CreatedAt, a2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	a := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      a.ID,
		Balance: util.RandomMoney(),
	}

	a2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, a2)

	require.Equal(t, a.ID, a2.ID)
	require.Equal(t, a.Owner, a2.Owner)
	require.Equal(t, arg.Balance, a2.Balance)
	require.Equal(t, a.Currency, a2.Currency)
	require.WithinDuration(t, a.CreatedAt, a2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	a := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), a.ID)
	require.NoError(t, err)

	a2, err := testQueries.GetAccount(context.Background(), a.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, a2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	a, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, a, 5)

	for _, a1 := range a {
		require.NotEmpty(t, a1)
	}
}
