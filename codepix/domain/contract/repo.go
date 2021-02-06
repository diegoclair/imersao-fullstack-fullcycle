package contract

import "github.com/diegoclair/imersao-fullstack-fullcycle/codepix/domain/model"

//RepoManager defines the repository aggregator interface
type RepoManager interface {
	Postgres() PostgresDB
}

//PostgresDB defines the postgres aggregator interface
type PostgresDB interface {
	Account() AccountRepo
}

// AccountRepo defines the data set for account repo
type AccountRepo interface {
	FindAccountByID(id string) (*model.Account, error)
}

// TransactionRepo defines the data set for account repo
type TransactionRepo interface {
	Register(transaction *model.Transaction)
}
