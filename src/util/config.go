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
func LoadConfig(path string, args ...interface{}) (config Config, err error) {
	name := "app"
	if len(args) >= 1 {
		name = args[0].(string)
	}
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
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
