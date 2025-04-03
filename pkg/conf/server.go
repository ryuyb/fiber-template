package conf

type ServerConfig struct {
	Addr string `yaml:"addr" env:"SERVER_ADDR" env-default:":8000"`
}
