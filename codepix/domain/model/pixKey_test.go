package model_test

import (
	"testing"

	"github.com/diegoclair/imersao-fullstack-fullcycle/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestModel_NewPixKey(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, _ := model.NewAccount(bank, accountNumber, ownerName)

	kind := model.PixKeyKindEmail
	key := "j@j.com"
	pixKey, err := model.NewPixKey(kind, account, key)

	require.NotEmpty(t, uuid.FromStringOrNil(pixKey.ID))
	require.Equal(t, pixKey.Kind, kind)
	require.Equal(t, pixKey.Status, "active")

	kind = model.PixKeyKindCPF
	key = "01234567890"
	_, err = model.NewPixKey(kind, account, key)
	require.Nil(t, err)

	_, err = model.NewPixKey("nome", account, key)
	require.NotNil(t, err)
}
