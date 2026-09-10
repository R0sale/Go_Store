package config

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string `mapstruct:"host"`
		Port     int    `mapstruct:"port"`
		User     string `mapstruct:"user"`
		Password string `mapstruct:"password"`
		DbName   string `mapstruct:"dbname"`
	} `mapstructure:"database"`

	SecretKey struct {
		Key string `mapstruct:"key"`
	} `mapstructure:"secretkey"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

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
