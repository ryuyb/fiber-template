package conf

type ServerConfig struct {
	Addr string `yaml:"addr" env:"SERVER_ADDR"`
}
