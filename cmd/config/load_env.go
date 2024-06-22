package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ALLOWED_ORIGINS string `mapstructure:"ALLOWED_ORIGINS"`
	PORT            string `mapstructure:"PORT"`
	AUTH_SERVICE    string `mapstructure:"AUTH_SERVICE"`
	VOCAB_SERVICE   string `mapstructure:"VOCAB_SERVICE"`
	GROUP_SERVICE   string `mapstructure:"GROUP_SERVICE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("broker")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
