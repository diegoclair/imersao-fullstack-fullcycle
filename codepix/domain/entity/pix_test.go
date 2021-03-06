package entity_test

import (
	"testing"

	"github.com/diegoclair/imersao/codepix/domain"
	"github.com/diegoclair/imersao/codepix/domain/entity"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestModel_NewPixKey(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := entity.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, _ := entity.NewAccount(bank, accountNumber, ownerName)

	keytype := domain.PixKeytypeEmail
	key := "j@j.com"
	pix, err := entity.NewPix(keytype, account, key)

	require.NotEmpty(t, uuid.FromStringOrNil(pix.ID))
	require.Equal(t, pix.KeyType, keytype)
	require.Equal(t, pix.Status, "active")

	keytype = domain.PixKeytypeCPF
	key = "01234567890"
	_, err = entity.NewPix(keytype, account, key)
	require.Nil(t, err)

	_, err = entity.NewPix("nome", account, key)
	require.NotNil(t, err)
}
