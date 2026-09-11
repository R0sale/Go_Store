package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstruct:"port"`
	} `mapstructure:"server"`

	Database struct {
		Host     string `mapstruct:"host"`
		Port     int    `mapstruct:"port"`
		User     string `mapstruct:"user"`
		Password string `mapstruct:"password"`
		DbName   string `mapstruct:"dbname"`
	} `mapstructure:"database"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(os.Getenv("CONFIG_PATH"))

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
