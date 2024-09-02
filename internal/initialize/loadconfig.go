package initialize

import (
	"fmt"

	"github.com/kaito-coder/go-ecommerce-architecture/global"
	"github.com/spf13/viper"
)


func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// Get value from config
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Print(err)
	}
}
