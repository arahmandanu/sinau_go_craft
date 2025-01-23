package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func GetRedisHost() string {
	return viper.GetString("REDIS_HOST")
}

func GetRedisPort() int {
	return viper.GetInt("REDIS_PORT")
}

func GetRedisUsername() string {
	return viper.GetString("REDIS_USERNAME")
}

func GetRedisPassword() string {
	return viper.GetString("REDIS_PASSWORD")
}

func GetRedisDB() int {
	return viper.GetInt("REDIS_DB")
}

func GetRedisDsn() string {
	return fmt.Sprintf("redis://%s:%s@%s:%d/%d", GetRedisUsername(), GetRedisPassword(), GetRedisHost(), GetRedisPort(), GetRedisDB())
}

func Init() {
	if os.Getenv("APP_ENV") == "development" {
		viper.SetConfigFile(".env")
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found")
		} else {
			fmt.Println("Config init error")
		}
	} else {
		println("Using config file:", viper.ConfigFileUsed())
	}

}
