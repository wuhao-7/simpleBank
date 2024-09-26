package db

import (
	"testing"
	"context"
	"time"
	"github.com/stretchr/testify/require"
	"github.com/wuhao-7/simplebank/util"
)

 func createRandomTransfers(t *testing.T,account1 Account,account2 Account) Transfer {
	 arg := CreateTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Amount: util.RandomMoney(),
	 }

	 transfer, err := testQueries.CreateTransfers(context.Background(), arg)
	
	 require.NoError(t,err)
	 require.NotEmpty(t,transfer)
		
	 require.Equal(t,arg.FromAccountID,transfer.FromAccountID)
	 require.Equal(t,arg.ToAccountID,transfer.ToAccountID)
	 require.Equal(t,arg.Amount,transfer.Amount)

	 require.NotZero(t,arg.Amount,transfer.Amount)
		
	 return transfer
 } 

 func TestCreateTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfers(t,account1,account2)

 }

 func TestUpdateTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	transfer1 := createRandomTransfers(t,account1,account2)

	require.NotEmpty(t,transfer1)

	arg :=UpdateTransfersParams{
		ID: transfer1.ID,
		Amount: util.RandomMoney(),
	}

	transfer2, err := testQueries.UpdateTransfers(context.Background(), arg)
	
	require.NoError(t,err)
	require.NotEmpty(t,transfer2)

	require.Equal(t,transfer2.FromAccountID,transfer1.FromAccountID)
	require.Equal(t,transfer2.ToAccountID,transfer1.ToAccountID)
	require.Equal(t,transfer2.Amount,arg.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)

 }

func TestDeleteTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	transfer1 := createRandomTransfers(t,account1,account2)

	require.NotEmpty(t,transfer1)
		

	err := testQueries.DeleteTransfers(context.Background(), transfer1.ID)

	require.NoError(t,err)
}

func TestListTransfers(t *testing.T){
	for i:=0; i<5 ;i++ {
		account1 := createRandomAccount(t)
		account2 := createRandomAccount(t)
		transfer1 := createRandomTransfers(t,account1,account2)
		require.NotEmpty(t,transfer1)
	}

	arg := ListTransfersParams{
		Limit: 5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(),arg)
	
	require.NoError(t,err)

	for _,transfer := range transfers {
		require.NotEmpty(t,transfer)
	}


}

