package conf

type LogConfig struct {
	Level string  `yaml:"level" env:"LOG_LEVEL" env-default:"INFO"`
	File  FileLog `yaml:"file"`
}

type FileLog struct {
	Enable     bool   `yaml:"enable" env:"LOG_FILE_ENABLE" env-default:"false"`
	Path       string `yaml:"path" env:"LOG_PATH" env-default:"./logs/app.log"`
	MaxSize    int    `yaml:"max_size" env:"LOG_MAX_SIZE" env-default:"10"`
	MaxBackups int    `yaml:"max_backups" env:"LOG_MAX_BACKUPS" env-default:"7"`
	MaxAge     int    `yaml:"max_age" env:"LOG_MAX_AGE" env-default:"30"`
	Compress   bool   `yaml:"compress" env:"LOG_COMPRESS" env-default:"true"`
}
