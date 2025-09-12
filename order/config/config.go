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
		PG   `yaml:"pg"`
	}

	App struct {
		Name    string `yaml:"name" env-required:"true" env:"APP_NAME"`
		Version string `yaml:"version" env-required:"true" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `yaml:"port" env-required:"true" env:"HTTP_PORT" env-default:"8080"`
	}

	PG struct {
		MaxPoolSize int    `yaml:"max_pool_size" env-required:"true" env:"PG_MAX_POOL_SIZE" env-default:"20"`
		URL         string `env-required:"true" env:"PG_URL"`
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
	fmt.Println(cfg)

	return cfg, nil
}
