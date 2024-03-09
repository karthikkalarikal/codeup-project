package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	GoSandboxServer string `mapstructure:"GO_SANDBOX_PORT"`
}

var envs = []string{
	"GO_SANDBOX_PORT",
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
