package config

import (
	"sync"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

//Environment standard values
const (
	EnvironmentTest string = "test"
)

// EnvironmentVariables is environment variables config
type EnvironmentVariables struct {
	Env      string
	Debug    bool
	Kafka    kafkaConfig
	Postgres postgresConfig
}

type kafkaConfig struct {
	BootstrapServers             string
	ConsumerGroupID              string
	TransactionTopic             string
	TransactionConfirmationTopic string
}

// postgresConfig is environment variables for postgress config
type postgresConfig struct {
	DSNTest     string
	DBTypeTest  string
	DBType      string
	DSN         string
	AutoMigrate bool
}

var (
	config *EnvironmentVariables
	once   sync.Once
)

// GetConfigEnvironment to read initial config
func GetConfigEnvironment() *EnvironmentVariables {
	once.Do(func() {
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		config = &EnvironmentVariables{}
		config.Env = cast.ToString(viper.Get("env"))
		config.Debug = cast.ToBool(viper.Get("debug"))
		config.Postgres.DSN = cast.ToString(viper.Get("dsnPostgres"))
		config.Postgres.DBType = cast.ToString(viper.Get("dbTypePostgres"))
		config.Postgres.DSNTest = cast.ToString(viper.Get("dsnTestPostgres"))
		config.Postgres.DBTypeTest = cast.ToString(viper.Get("dbTypeTestPostgres"))
		config.Postgres.AutoMigrate = cast.ToBool(viper.Get("autoMigrateDBPostgres"))
		config.Kafka.BootstrapServers = cast.ToString(viper.Get("kafkaBootstrapServers"))
		config.Kafka.ConsumerGroupID = cast.ToString(viper.Get("kafkaConsumerGroupID"))
		config.Kafka.TransactionTopic = cast.ToString(viper.Get("kafkaTransactionTopic"))
		config.Kafka.TransactionConfirmationTopic = cast.ToString(viper.Get("kafkaTransactionConfirmationTopic"))
	})
	return config
}
