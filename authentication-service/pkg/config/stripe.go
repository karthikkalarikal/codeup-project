package config

import "github.com/spf13/viper"

func NewStripeConfig() (config *StripeConfig, err error) {

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	for _, env := range stripeenvs {
		if err = viper.BindEnv(env); err != nil {
			return
		}
	}
	err = viper.Unmarshal(&config)

	return
}

type StripeConfig struct {
	StripeKey    string `mapstructure:"STRIPE_PUB_KEY"`
	StripeSecret string `mapstructure:"STRIPE_SECRET_KEY"`
}

var stripeenvs = []string{
	"STRIPE_SECRET_KEY", "STRIPE_PUB_KEY",
}
