package entity_test

import (
	"testing"

	"github.com/diegoclair/imersao/codepix/domain/entity"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewTransaction(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := entity.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, _ := entity.NewAccount(bank, accountNumber, ownerName)

	accountNumberDestination := "abcdestination"
	ownerName = "Mariana"
	accountDestination, _ := entity.NewAccount(bank, accountNumberDestination, ownerName)

	keyType := entity.PixKeytypeEmail
	key := "j@j.com"
	pix, _ := entity.NewPix(keyType, accountDestination, key)

	require.NotEqual(t, account.ID, accountDestination.ID)

	amount := 3.10
	statusTransaction := "pending"

	transaction, err := entity.NewTransaction(account, amount, pix, "My description", "")

	require.Nil(t, err)
	require.NotNil(t, uuid.FromStringOrNil(transaction.ID))
	require.Equal(t, transaction.Amount, amount)
	require.Equal(t, transaction.Status, statusTransaction)
	require.Equal(t, transaction.Description, "My description")
	require.Empty(t, transaction.CancelDescription)

	pixSameAccount, err := entity.NewPix(keyType, account, key)

	_, err = entity.NewTransaction(account, amount, pixSameAccount, "My description", "")
	require.NotNil(t, err)

	_, err = entity.NewTransaction(account, 0, pix, "My description", "")
	require.NotNil(t, err)

}

func TestModel_ChangeStatusOfATransaction(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := entity.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, _ := entity.NewAccount(bank, accountNumber, ownerName)

	accountNumberDestination := "abcdestination"
	ownerName = "Mariana"
	accountDestination, _ := entity.NewAccount(bank, accountNumberDestination, ownerName)

	kind := "email"
	key := "j@j.com"
	pix, _ := entity.NewPix(kind, accountDestination, key)

	amount := 3.10
	transaction, _ := entity.NewTransaction(account, amount, pix, "My description", "")

	transaction.Complete()
	require.Equal(t, transaction.Status, entity.TransactionCompleted)

	transaction.Cancel("Error")
	require.Equal(t, transaction.Status, entity.TransactionError)
	require.Equal(t, transaction.CancelDescription, "Error")
}
