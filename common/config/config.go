package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Conf *Config

type ApiConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
type MongoConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type MinioConfig struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	AccessKeyID     string `json:"access_key_id" yaml:"access_key_id" mapstructure:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key" yaml:"secret_access_key" mapstructure:"secret_access_key"`
	UseSSL          bool   `json:"use_ssl" yaml:"use_ssl" mapstructure:"use_ssl"`
}
type Config struct {
	Mongo MongoConfig `json:"mongo"`
	Api   ApiConfig   `json:"api"`
	Minio MinioConfig `json:"minio"`
}

func ReadConfig() {
	viper.SetConfigName("config_pro")
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
