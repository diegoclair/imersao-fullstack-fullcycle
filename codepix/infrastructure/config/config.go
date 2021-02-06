package config

import (
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

// EnvironmentConfig is environment variables config
type EnvironmentConfig struct {
	Env      string
	Debug    bool
	Postgres PostgresConfig
}

// PostgresConfig is environment variables for postgress config
type PostgresConfig struct {
	DSNTest     string
	DBTypeTest  string
	DBType      string
	DSN         string
	AutoMigrate bool
}

// GetConfig to read initial config
func GetConfig() (config EnvironmentConfig) {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	config.Env = cast.ToString(viper.Get("env"))
	config.Debug = cast.ToBool(viper.Get("debug"))
	config.Postgres.DSN = cast.ToString(viper.Get("dsnPostgres"))
	config.Postgres.DBType = cast.ToString(viper.Get("dbTypePostgres"))
	config.Postgres.DSNTest = cast.ToString(viper.Get("dsnTestPostgres"))
	config.Postgres.DBTypeTest = cast.ToString(viper.Get("dbTypeTestPostgres"))

	return
}
