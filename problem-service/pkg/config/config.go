package config

import "github.com/spf13/viper"

type Config struct {
	ProblemServicePort string `mapstructure:"PROBLEM_SERVICE_PORT"`
	MongoURL           string `mapstructuure:"MONGO_URL"`
	Username           string `mapstructuure:"MONGO_INITDB_DATABASE"`
	Password           string `mapstructuure:"MONGO_INITDB_ROOT_USERNAME"`
	AuthSource         string `mapstructuure:"MONGO_INITDB_ROOT_PASSWORD"`
	AuthMechanism      string `mapstructuure:"AUTH_MECHANISM"`
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
			return // naked return
		}
	}
	err = viper.Unmarshal(&config)
	return
}
