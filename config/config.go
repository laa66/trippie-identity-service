package config

import (
	"os"

	apperr "github.com/laa66/trippie-identity-service.git/error"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/logger"
	"gopkg.in/yaml.v3"
)

var appConfig *Config

type (
	Config struct {
		HTTP HTTP `yaml:"http"`
		DB   DB   `yaml:"db"`
	}

	HTTP struct {
		Port int `yaml:"port"`
	}

	DB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	}
)

func LoadConfig(path string) (*apperr.AppErr) {
	data, err := os.ReadFile(path)
	if err != nil {
		logger.Log().Error("config file read", "error", err)
		return apperr.Wrap(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		logger.Log().Error("yaml file unmarshal", "error", err)
		return apperr.Wrap(err)
	}

	appConfig = &cfg
	return nil
}


func GetConfig() *Config {
	return appConfig
}