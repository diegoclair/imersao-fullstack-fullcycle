package contract

import "github.com/diegoclair/imersao/codepix/domain/entity"

//PostgresRepo defines the postgres aggregator interface
type PostgresRepo interface {
	Account() AccountRepo
	Bank() BankRepo
	Pix() PixRepo
	Transaction() TransactionRepo
	Close() error
}

// AccountRepo defines the data set for account repo
type AccountRepo interface {
	AddAccount(account *entity.Account) error
	FindAccountByID(id string) (*entity.Account, error)
}

// BankRepo defines the data set for Bank repo
type BankRepo interface {
	AddBank(bank *entity.Bank) error
	FindBankByID(id string) (*entity.Bank, error)
}

// PixRepo defines the data set for Pix repo
type PixRepo interface {
	AddPixKey(pixKey *entity.Pix) (*entity.Pix, error)
	FindPixByKey(key, kind string) (*entity.Pix, error)
}

// TransactionRepo defines the data set for transaction repo
type TransactionRepo interface {
	FindByID(id string) (*entity.Transaction, error)
	Register(transaction *entity.Transaction) error
	Save(transaction *entity.Transaction) error
}
