package config

import (
	"github.com/go-playground/validator/v10"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/spf13/viper"
)

type Config struct {
	Port              string `mapstructure:"PORT"`
	AuthServiceUrl    string `mapstructure:"AUTH_SERVICE_URL"`
	ProblemServiceUrl string `mapstructure:"PROBLEM_SERVICE_URL"`
	GoSandboxUrl      string `mapstructure:"GO_SANDBOX_PORTAL"`
	JWT               *echojwt.Config
}

func NewConfig() (*Config, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

var envs = []string{"PORT", "AUTH_SERVICE_URL", "GO_SANDBOX_PORTAL", "PROBLEM_SERVICE_URL"}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile("api-gateway.env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}
