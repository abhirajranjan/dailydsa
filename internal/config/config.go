package config

import (
	"github.com/spf13/viper"
)

var Config *config

func LoadConfig(path string) (config *config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("cfg")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	Config = config
	return
}
