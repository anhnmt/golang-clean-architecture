package config

type Log struct {
	Format string `yaml:"format" env:"LOG_FORMAT" env-default:"text"`
	Level  string `yaml:"level" env:"LOG_LEVEL" env-default:"info"`

	// Log file
	File       string `yaml:"file" env:"LOG_FILE"`
	MaxSize    int    `yaml:"max_size" env:"LOG_MAX_SIZE" env-default:"100"` // MB
	MaxBackups int    `yaml:"max_backups" env:"LOG_MAX_BACKUPS" env-default:"5"`
	MaxAge     int    `yaml:"max_age" env:"LOG_MAX_AGE" env-default:"28"` // days
}
