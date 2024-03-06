package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ProblemServicePort string `mapstructure:"PROBLEM_SERVICE_PORT"`
	MongoURL           string `mapstructure:"MONGO_URL"`
	MongoDataBase      string `mapstructure:"MONGO_INITDB_DATABASE"`
	Username           string `mapstructure:"MONGO_INITDB_ROOT_USERNAME"`
	Password           string `mapstructure:"MONGO_INITDB_ROOT_PASSWORD"`
	AuthMechanism      string `mapstructure:"AUTH_MECHANISM"`
}

var envs = []string{
	"PROBLEM_SERVICE_PORT",
	"MONGO_URL",
	"MONGO_INITDB_DATABASE",
	"MONGO_INITDB_ROOT_USERNAME",
	"MONGO_INITDB_ROOT_PASSWORD",
	"AUTH_MECHANISM",
}

func LoadConfig() (config *Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	for _, env := range envs {
		if err = viper.BindEnv(env); err != nil {
			fmt.Println("here")
			fmt.Println("err", err)
			return // naked return
		}
	}
	err = viper.Unmarshal(&config)
	return
}
