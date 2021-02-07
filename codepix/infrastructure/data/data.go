package data

import (
	"log"

	"github.com/diegoclair/imersao/codepix/domain/contract"
	"github.com/diegoclair/imersao/codepix/infrastructure/data/postgres"
)

// func init() {
// 	_, b, _, _ := runtime.Caller(0)
// 	basepath := filepath.Dir(b)

// 	err := godotenv.Load(basepath + "/../../.env")
// 	if err != nil {
// 		log.Fatal("Error loading .env files")
// 	}
// }

//we can add here more than one database
type data struct {
	postgresRepo contract.PostgresRepo
}

// Connect returns a instace of postgres db
func Connect() (contract.DataManager, error) {
	repo := new(data)
	return &data{
		postgresRepo: repo.Postgres(),
	}, nil
}

func (d *data) Postgres() contract.PostgresRepo {
	postgresRepo, err := postgres.Instance()
	if err != nil {
		log.Fatalf("Error to start postgres instance: %v", err)
	}
	return postgresRepo
}
