package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/diegoclair/imersao-fullstack-fullcycle/domain/model"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	_ "gorm.io/driver/sqlite"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.env")
	if err != nil {
		log.Fatal("Error loading .env files")
	}
}

//ConnectDB to create a connection with database
func ConnectDB(env string) *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	if env == "test" {
		dsn = os.Getenv("dsnTest")
		db, err = gorm.Open(os.Getenv("dbTypeTest"), dsn)

	} else {
		dsn = os.Getenv("dsn")
		db, err = gorm.Open(os.Getenv("dbType"), dsn)
	}
	if err != nil {
		log.Fatalf("Error to connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv("debug") == "true" {
		db.LogMode(true)
	}

	if os.Getenv("AutoMigrationDB") == "true" {
		db.AutoMigrate(&model.Bank{}, &model.Account{}, &model.PixKey{}, &model.Transaction{})
	}

	return db

}
