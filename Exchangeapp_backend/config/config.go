package config

import (
	"log"
	"github.com/spf13/viper"
)

// 配置 struct体
type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	
}

var Appconfig *Config

func Initconfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %v", err)
	}

	Appconfig =&Config{}


    if err := viper.Unmarshal(Appconfig); err != nil {
		log.Fatalf("Unable decode to config struct, %v", err)
	}

}

