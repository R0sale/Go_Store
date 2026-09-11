package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
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
		return nil, fmt.Errorf("couldn't load env variables")
	}

	configPath := os.Getenv("CONFIG_PATH")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("couldn't read the config file at %s %v", configPath, err.Error())
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal the config into struct")
	}

	return &cfg, nil
}
