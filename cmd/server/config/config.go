package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/anhnmt/golang-clean-architecture/pkg/config"
)

type Config struct {
	config.App `yaml:"app"`
	config.Log `yaml:"log"`
}

func New() (*Config, error) {
	cfg := &Config{}

	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("getwd error: %w", err)
	}

	path := fmt.Sprintf("%s/%s", dir, "config.yml")
	err = cleanenv.ReadConfig(filepath.ToSlash(path), cfg)
	if err != nil {
		return nil, fmt.Errorf("read config error: %w", err)
	}

	return cfg, nil
}
