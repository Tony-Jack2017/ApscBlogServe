package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Conf *Config

type MongoConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ApiConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Config struct {
	Mongo MongoConfig `json:"mongo"`
	Api   ApiConfig   `json:"api"`
}

func ReadConfig() {
	viper.SetConfigName("config_dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./common/config/yaml")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println(err)
			return
		} else {
			// Config file was found but another error was produced\
			fmt.Println(err)
			return
		}
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		fmt.Println("Unmarshal the config struct to 'Config' failed !!!")
		return
	}
	fmt.Println("Reading the config file successful.")
}
