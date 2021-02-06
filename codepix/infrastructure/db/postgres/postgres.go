package db

import (
	"log"

	"github.com/diegoclair/imersao-fullstack-fullcycle/codepix/domain/contract"
	"github.com/diegoclair/imersao-fullstack-fullcycle/codepix/domain/model"
	"github.com/diegoclair/imersao-fullstack-fullcycle/codepix/infrastructure/config"
	"github.com/diegoclair/imersao-fullstack-fullcycle/codepix/infrastructure/db/postgres"
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
	_ "gorm.io/driver/sqlite"
)

// dbManager is the Database connection manager
type dbManager struct {
	db *gorm.DB
}

//Instance to create a connection with database
func Instance(cfg config.EnvironmentConfig) (contract.RepoManager, error) {
	var dsn string
	var db *gorm.DB
	var err error

	if cfg.Env == "test" {
		dsn = cfg.Postgres.DSNTest
		dbTypeTest := cfg.Postgres.DBTypeTest
		db, err = gorm.Open(dbTypeTest, dsn)

	} else {
		dsn = cfg.Postgres.DSN
		dbType := cfg.Postgres.DBType
		db, err = gorm.Open(dbType, dsn)
	}
	if err != nil {
		log.Fatalf("Error to connecting to database: %v", err)
		return nil, err
	}

	db.LogMode(cfg.Debug)

	if cfg.Postgres.AutoMigrate {
		db.AutoMigrate(&model.Bank{}, &model.Account{}, &model.Pix{}, &model.Transaction{})
	}

	connection := &dbManager{
		db: db,
	}
	return connection, nil
}

//Account returns the account set
func (c *dbManager) Account() contract.AccountRepo {
	return postgres.NewAccountRepo(c.db)
}

//Bank returns the bank set
func (c *dbManager) Bank() contract.BankRepo {
	return postgres.NewBankRepo(c.db)
}

//Pix returns the pix set
func (c *dbManager) Pix() contract.PixRepo {
	return postgres.NewPixRepo(c.db)
}

//Transaction returns the transaction set
func (c *dbManager) Transaction() contract.TransactionRepo {
	return postgres.NewTransactionRepo(c.db)
}
