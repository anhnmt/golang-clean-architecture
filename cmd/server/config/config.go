package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/anhnmt/golang-clean-architecture/pkg/config"
)

type Config struct {
	App      config.App      `yaml:"app" json:"app"`
	Log      config.Log      `yaml:"log" json:"log"`
	Postgres config.Postgres `yaml:"postgres" json:"postgres"`
	Server   config.Server   `yaml:"server" json:"server"`
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
