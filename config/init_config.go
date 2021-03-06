package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("toml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}
