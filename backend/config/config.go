package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Http     Http     `mapstructure:"http"`
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
}

type Http struct {
	Timeout int `mapstructure:"timeout"`
}

type Server struct {
	Port string `mapstructure:"port"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
}

func GetConfig() Config {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	// config.Database.Host = os.Getenv("MYSQL_HOST")
	// config.Database.Port = os.Getenv("MYSQL_PORT")

	return config
}
