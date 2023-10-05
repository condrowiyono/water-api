package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfiguration struct {
	Driver   string
	Dbname   string
	Username string
	Password string
	Host     string
	Port     string
	LogMode  bool
}

func DbConfiguration() string {
	masterDBName := viper.GetString("MASTER_DB_NAME")
	masterDBUser := viper.GetString("MASTER_DB_USER")
	masterDBPassword := viper.GetString("MASTER_DB_PASSWORD")
	masterDBHost := viper.GetString("MASTER_DB_HOST")
	masterDBPort := viper.GetString("MASTER_DB_PORT")

	masterDBDSN := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		masterDBUser, masterDBPassword, masterDBHost, masterDBPort, masterDBName,
	)

	return masterDBDSN
}
