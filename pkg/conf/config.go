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

var cfg *AppConfig

type AppConfig struct {
	Environment Env            `yaml:"env" env:"APP_ENV" env-default:"prod"` // prod, dev
	Server      ServerConfig   `yaml:"server"`
	Log         LogConfig      `yaml:"log"`
	Database    DatabaseConfig `yaml:"database"`
	Jwt         JwtConfig      `yaml:"jwt"`
}

func (c *AppConfig) IsProdEnv() bool {
	return c.Environment == ProdEnv
}

func (c *AppConfig) IsDevEnv() bool {
	return c.Environment == DevEnv
}

func Initialize(path string) AppConfig {
	if cfg != nil {
		log.Warnf("AppConfig is already initialized, replace it with new config path: %s", path)
	}
	cfg = &AppConfig{}
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		log.Fatalf("read configuration failed, path: %s, error: %v", path, err)
	}
	log.Debugf("Configuration loaded from path: %s", path)
	return *cfg
}

func GetConfig() AppConfig {
	return *cfg
}
