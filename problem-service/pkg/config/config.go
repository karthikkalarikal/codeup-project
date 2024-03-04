package config

import "github.com/spf13/viper"

type Config struct {
	ProblemServicePort string `mapstructure:"PROBLEM_SERVICE_PORT"`
}

var envs = []string{
	"PROBLEM_SERVICE_PORT",
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
