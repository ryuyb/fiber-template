package conf

type DatabaseConfig struct {
	Driver  string `yaml:"driver" env:"DB_DRIVER"`
	Source  string `yaml:"source" env:"DB_SOURCE"`
	Migrate bool   `yaml:"migrate" env:"DB_MIGRATE"`
}
