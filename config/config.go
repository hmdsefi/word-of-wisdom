package config

import (
	"fmt"
	"os"
	"path"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type ServerConfig struct {
	common
	Port   string `yaml:"port" env:"PORT"`
	APIKey string `yaml:"server_address" env:"API_KEY"`
}

type ClientConfig struct {
	common
	ServerAddress string `yaml:"server_address" env:"SERVER_ADDRESS"`
}

type common struct {
	SecretKey string `yaml:"secret_key" env:"SECRET_KEY"`
}

// newConfig returns app config.
func newConfig(cfg any) (any, error) {
	projectRoot, err := os.Getwd()
	envPath := path.Join(projectRoot, ".env")

	if err != nil {
		return nil, fmt.Errorf("path.Join: %w", err)
	}

	_, err = os.Stat(envPath)
	if err == nil {
		err = godotenv.Load(envPath)
		if err != nil {
			return nil, fmt.Errorf("godotenv.Load: %w", err)
		}
	} else {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf(".env reading error: %w", err)
		}
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func NewClientConfig() (*ClientConfig, error) {
	cfg, err := newConfig(&ClientConfig{})
	if err != nil {
		return nil, err
	}

	cliCfg, _ := cfg.(*ClientConfig)
	return cliCfg, nil
}

func NewServerConfig() (*ServerConfig, error) {
	cfg, err := newConfig(&ServerConfig{})
	if err != nil {
		return nil, err
	}

	svrCfg, _ := cfg.(*ServerConfig)
	return svrCfg, nil
}
