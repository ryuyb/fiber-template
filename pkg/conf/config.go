package conf

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/ilyakaznacheev/cleanenv"
)

type Env string

const (
	DevEnv  = "dev"
	ProdEnv = "prod"
)

type AppConfig struct {
	Environment Env            `yaml:"env" env:"APP_ENV" env-default:"prod"` // prod, dev
	Server      ServerConfig   `yaml:"server"`
	Log         LogConfig      `yaml:"log"`
	Database    DatabaseConfig `yaml:"database"`
}

func (c *AppConfig) IsProdEnv() bool {
	return c.Environment == ProdEnv
}

func (c *AppConfig) IsDevEnv() bool {
	return c.Environment == DevEnv
}

func New(path string) AppConfig {
	var cfg AppConfig

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatalf("read configuration failed: %v", err)
	}

	return cfg
}
