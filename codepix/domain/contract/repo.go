package contract

import "github.com/diegoclair/imersao/codepix/domain/model"

//PostgresRepo defines the postgres aggregator interface
type PostgresRepo interface {
	Account() AccountRepo
	Bank() BankRepo
	Pix() PixRepo
	Transaction() TransactionRepo
}

// AccountRepo defines the data set for account repo
type AccountRepo interface {
	AddAccount(account *model.Account) error
	FindAccountByID(id string) (*model.Account, error)
}

// BankRepo defines the data set for Bank repo
type BankRepo interface {
	AddBank(bank *model.Bank) error
	FindBankByID(id string) (*model.Bank, error)
}

// PixRepo defines the data set for Pix repo
type PixRepo interface {
	AddPixKey(pixKey *model.Pix) (*model.Pix, error)
	FindPixKeyByID(key, kind string) (*model.Pix, error)
}

// TransactionRepo defines the data set for transaction repo
type TransactionRepo interface {
	FindByID(id string) (*model.Transaction, error)
	Register(transaction *model.Transaction) error
	Save(transaction *model.Transaction) error
}
