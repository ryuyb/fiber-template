package conf

type JwtConfig struct {
	SigningKey string `yaml:"signing_key" env:"JWT_SIGNING_KEY"`
}
