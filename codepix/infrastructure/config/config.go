package config

import (
	"sync"

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
		viper.AutomaticEnv() //to read from environment variables and it will override values that it has read from .env file with the values of the corresponding environment variables if they exist

		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}

		config = &EnvironmentVariables{}
		config.Env = viper.GetString("ENV")
		config.Debug = viper.GetBool("DEBUG")
		config.Postgres.DSN = viper.GetString("DSN_POSTGRES")
		config.Postgres.DBType = viper.GetString("DB_TYPE_POSTGRES")
		config.Postgres.DSNTest = viper.GetString("DSN_TEST_POSTGRES")
		config.Postgres.DBTypeTest = viper.GetString("DB_TYPE_TEST_POSTGRES")
		config.Postgres.AutoMigrate = viper.GetBool("AUTO_MIGRATE_DB_POSTGRES")
		config.Kafka.BootstrapServers = viper.GetString("KAFKA_BOOTSTRAP_SERVERS")
		config.Kafka.ConsumerGroupID = viper.GetString("KAFKA_CONSUMER_GROUP_ID")
		config.Kafka.TransactionTopic = viper.GetString("KAFKA_TRANSACTION_TOPIC")
		config.Kafka.TransactionConfirmationTopic = viper.GetString("KAFKA_TRANSACTION_CONFIRMATIONTOPIC")

	})
	return config
}
