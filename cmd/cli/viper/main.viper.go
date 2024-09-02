package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Database []struct {
		User     string `mapstructure:"user"`
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Password string `mapstructure:"password"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// Get value from config
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Print(err)
	}
	fmt.Println(config.Server.Port)
	for _, db := range config.Database {
		fmt.Printf("User: %s, Host: %s, Port: %s, Password: %s\n", db.User, db.Host, db.Port, db.Password)
	}
}
