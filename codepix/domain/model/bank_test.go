package model

import (
	"testing"

	"gopkg.in/stretchr/testify.v1/require"
)

func TestNewBank(t *testing.T) {

	code := "001"
	name := "Banco do Brasil"

	bank, err := NewBank(code, name)

	require.Nil(t, err)
	require.NotEmpty(t, bank.ID)
	require.Equal(t, bank.Code, code)
	require.Equal(t, bank.Name, name)

	_, err = NewBank(code, "")
	require.NotNil(t, err)

	_, err = NewBank("", name)
	require.NotNil(t, err)

	_, err = NewBank("", "")
	require.NotNil(t, err)
}
