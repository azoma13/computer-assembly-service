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
		HTTP `yaml:"http"`
		Pg   `yaml:"pg"`
		Grpc `yaml:"grpc"`
	}

	App struct {
		Name    string `yaml:"name" env-required:"true" env:"APP_NAME"`
		Version string `yaml:"version" env-required:"true" env:"APP_VERSION"`
	}

	HTTP struct {
		Host string `yaml:"host" env-required:"true" env:"HTTP_HOST"`
		Port string `yaml:"port" env-required:"true" env:"HTTP_PORT"`
	}

	Pg struct {
		URL string `yaml:"url" env-required:"true" env:"PG_URL"`

		MaxPoolSize int `yaml:"max_pool_size" env-required:"true" env:"PG_MAX_POOL_SIZE" env-default:"20"`
	}

	Grpc struct {
		Hardware string `yaml:"hardware" env-required:"true" env:"GRPC_HARDWARE"`

		Payment string `yaml:"payment" env-required:"true" env:"GRPC_PAYMENT"`
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
