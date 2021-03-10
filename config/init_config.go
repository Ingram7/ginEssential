package config

import (
	"os"

	"github.com/spf13/viper"
)

func InitConfig() error {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("toml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil

}
