package contract

import "github.com/diegoclair/imersao/codepix/domain/entity"

// PixService holds a pix service operations
type PixService interface {
	FindKeyByID(key, kind string) (*entity.Pix, error)
	RegisterKey(key, kind, accountID string) (*entity.Pix, error)
}

// TransactionService holds a transaction service operations
type TransactionService interface {
	Register(accountIDFrom, pixKeyTo, pixKeyKindTo, description, id string, amount float64) (*entity.Transaction, error)
	Confirm(transactionID string) (*entity.Transaction, error)
	Complete(transactionID string) (*entity.Transaction, error)
	Error(transactionID, reason string) (*entity.Transaction, error)
}
