package db

import (
	"context"
	"testing"
	"github.com/wuhao-7/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntries(t *testing.T, account Account) Entry {
	arg := CreateEntriesParams{
		AccountID: account.ID,
		Amount: util.RandomMoney(),
	}
	
	entry,err := testQueries.CreateEntries(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t,arg.AccountID,entry.AccountID)
	require.Equal(t,arg.Amount,entry.Amount)

	require.NotZero(t,entry.ID)
	require.NotZero(t,entry.CreatedAt)

	return entry
}
 func TestCreateEntries(t *testing.T) {
	account:= createRandomAccount(t)
	createRandomEntries(t,account)

 }

 func TestGetEntry(t *testing.T){

	account := createRandomAccount(t)

	entry1 := createRandomEntries(t,account) 
	entry2, err := testQueries.GetEntries(context.Background(), entry1.ID)
	
	require.NoError(t,err)
	require.NotEmpty(t,entry2)

	require.Equal(t,entry1.ID,entry2.ID)
	require.Equal(t,entry1.AccountID,entry2.AccountID)
	require.Equal(t,entry1.Amount,entry2.Amount)
	
 }

 func TestUpdateEntry(t *testing.T) {
	account1 := createRandomAccount(t)
	arg := UpdateEntriesParams{
		AccountID: account1.ID,
		Amount: util.RandomMoney(),
	}

	entry1 := createRandomEntries(t,account1) 

	entry2,err := testQueries.UpdateEntries(context.Background(),arg)

	require.NoError(t,err)
	require.NotEmpty(t,entry1)
	require.NotEmpty(t,entry2)

	require.Equal(t,entry1.ID,entry2.ID)
	require.Equal(t,entry1.AccountID,entry2.AccountID)
	require.Equal(t,arg.Amount,entry2.Amount)



 }
 func TestDeleteEntry(t *testing.T) {
	account1 := createRandomAccount(t)
	entry1 := createRandomEntries(t,account1)

	err := testQueries.DeleteEntries(context.Background(),entry1.ID)

	require.NoError(t,err)
	


 }

 func TestListEntries(t *testing.T) {
	for i:=0; i<10; i++ {
		account := createRandomAccount(t)
		createRandomEntries(t,account)
	}

	arg := ListEntriesParams{
		Limit: 5,
		Offset: 5,
	}
	
	entries,err := testQueries.ListEntries(context.Background(),arg)

	require.NoError(t,err)
	require.Len(t,entries,5)
	
	for _,entry := range entries {
		require.NotEmpty(t,entry)
	}

 }





