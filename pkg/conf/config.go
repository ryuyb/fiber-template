package conf

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

func New(path string) AppConfig {
	var cfg AppConfig

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatalf("read configuration failed: %v", err)
	}

	return cfg
}
