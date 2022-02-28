package util

import (
	"github.com/spf13/viper"
)

// Config config variable loaded from environment
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBName        string `mapstructure:"DB_NAME"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	SecretKey     string `mapstructure:"SECRET_KEY"`
}

// LoadConfig load config variable from env file
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	return
}
