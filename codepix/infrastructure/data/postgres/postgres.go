package postgres

import (
	"sync"

	"github.com/diegoclair/imersao/codepix/contract"
	"github.com/diegoclair/imersao/codepix/domain/entity"
	"github.com/diegoclair/imersao/codepix/util/config"
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq" //is a pure Go Postgres driver for the database/sql package
	_ "gorm.io/driver/sqlite"
)

var (
	connection *postgres
	onceDB     sync.Once
	connErr    error
)

// postgres is the Database connection manager
type postgres struct {
	db *gorm.DB
}

//Instance to create a connection with database
func Instance() (contract.PostgresRepo, error) {
	onceDB.Do(func() {
		cfg := config.GetConfigEnvironment()

		var dsn, dbType string
		var db *gorm.DB
		var err error

		dsn = cfg.Postgres.DSN
		dbType = cfg.Postgres.DBType

		if cfg.Env == config.EnvironmentTest {
			dsn = cfg.Postgres.DSNTest
			dbType = cfg.Postgres.DBTypeTest
		}

		db, err = gorm.Open(dbType, dsn)
		if err != nil {
			connErr = err
		}

		db.LogMode(cfg.Debug)

		if cfg.Postgres.AutoMigrate {
			db.AutoMigrate(&entity.Bank{}, &entity.Account{}, &entity.Pix{}, &entity.Transaction{})
		}

		connection = &postgres{
			db: db,
		}
	})
	return connection, connErr
}

// Close closes the db connection
func (c *postgres) Close() (err error) {
	return c.db.Close()
}

//Account returns the account set
func (c *postgres) Account() contract.AccountRepo {
	return newAccountRepo(c.db)
}

//Bank returns the bank set
func (c *postgres) Bank() contract.BankRepo {
	return newBankRepo(c.db)
}

//Pix returns the pix set
func (c *postgres) Pix() contract.PixRepo {
	return newPixRepo(c.db)
}

//Transaction returns the transaction set
func (c *postgres) Transaction() contract.TransactionRepo {
	return newTransactionRepo(c.db)
}
