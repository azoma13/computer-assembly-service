package config

import (
	"fmt"
	"path"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		App  `yaml:"app"`
		Grpc `yaml:"grpc"`
	}

	App struct {
		Name    string `yaml:"name" env-required:"true" env:"APP_NAME"`
		Version string `yaml:"version" env-required:"true" env:"APP_VERSION"`
	}

	Grpc struct {
		Port string `yaml:"port" env-required:"true" env:"GRPC_PORT"`
		Host string `yaml:"host" env-required:"host" env:"GRPC_HOST"`
		Addr string `yaml:"addr" env-required:"true" env:"GRPC_ADDR"`
	}
)

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}

	err := godotenv.Load()
	if err != nil {
		return &Config{}, fmt.Errorf("unable to load .env file: %w", err)
	}

	err = cleanenv.ReadConfig(path.Join("./", configPath), cfg)
	if err != nil {
		return &Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	return cfg, nil
}
