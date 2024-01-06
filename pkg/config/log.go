package config

type Log struct {
	Format string `yaml:"format" env:"LOG_FORMAT" env-default:"text"`
	Level  string `yaml:"level" env:"LOG_LEVEL" env-default:"info"`
}
