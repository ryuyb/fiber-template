package conf

type JwtConfig struct {
	SigningKey         string `yaml:"signing_key" env:"JWT_SIGNING_KEY"`
	ValidWithinMinutes uint32 `yaml:"valid_within_minutes" env:"JWT_VALID_WITHIN_MINUTES"`
}
