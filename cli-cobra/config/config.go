package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error: %v", err)
		panic(err)
	}
}
