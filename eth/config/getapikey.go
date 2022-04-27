package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetAPIKey() (string, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	return viper.GetString("APIKEY"), nil
}
