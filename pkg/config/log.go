package config

type Log struct {
	Format string `yaml:"format" json:"format,omitempty" env:"LOG_FORMAT" env-default:"text"`
	Level  string `yaml:"level" json:"level,omitempty" env:"LOG_LEVEL" env-default:"info"`

	// Log file
	File       string `yaml:"file" json:"file,omitempty" env:"LOG_FILE"`
	MaxSize    int    `yaml:"max_size" json:"max_size,omitempty" env:"LOG_MAX_SIZE" env-default:"100"` // MB
	MaxBackups int    `yaml:"max_backups" json:"max_backups,omitempty" env:"LOG_MAX_BACKUPS" env-default:"5"`
	MaxAge     int    `yaml:"max_age" json:"max_age,omitempty" env:"LOG_MAX_AGE" env-default:"28"` // days
}
