package conf

type DatabaseConfig struct {
	Driver  string `yaml:"driver" env:"DB_DRIVER" env-default:"sqlite3"`
	Source  string `yaml:"source" env:"DB_SOURCE" env-default:":memory:?_fk=1&_pragma=foreign_keys(1)"`
	Migrate bool   `yaml:"migrate" env:"DB_MIGRATE" env-default:"true"`
}
