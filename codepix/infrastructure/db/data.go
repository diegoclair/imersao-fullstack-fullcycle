package db

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/diegoclair/imersao/codepix/domain/contract"
	"github.com/diegoclair/imersao/codepix/infrastructure/config"
	"github.com/diegoclair/imersao/codepix/infrastructure/db/postgres"
	"github.com/joho/godotenv"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.env")
	if err != nil {
		log.Fatal("Error loading .env files")
	}
}

// Connect returns a instace of postgres db
func Connect(cfg config.EnvironmentConfig) (contract.RepoManager, error) {
	return postgres.Instance(cfg)
}
