package entity_test

import (
	"testing"

	"github.com/diegoclair/imersao/codepix/domain/entity"
	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/require"
)

func TestModel_NewAccount(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := entity.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Diego"
	account, err := entity.NewAccount(bank, ownerName, accountNumber)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(account.ID))
	require.Equal(t, account.Number, accountNumber)
	require.Equal(t, account.OwnerName, ownerName)
	//require.Equal(t, account.Bank.ID, bank.ID)

	_, err = entity.NewAccount(bank, "", ownerName)
	require.NotNil(t, err)
}
