package config

import (
	"fmt"
	"path"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		App   `yaml:"app"`
		Grpc  `yaml:"grpc"`
		Mongo `yaml:"mongo"`
	}

	App struct {
		Name    string `yaml:"name" env-required:"true" env:"APP_NAME"`
		Version string `yaml:"version" env-required:"true" env:"APP_VERSION"`
	}

	Grpc struct {
		Addr string `yaml:"addr" env-required:"true" env:"GRPC_ADDR"`
	}

	Mongo struct {
		Url string `yaml:"url" env-required:"true" env:"MONGO_URL"`

		MaxPoolSize int `yaml:"max_pool_size" env-required:"true" env:"MONGO_MAX_POOL_SIZE"`
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
