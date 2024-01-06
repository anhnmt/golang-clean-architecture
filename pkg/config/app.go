package config

type App struct {
	Name string `env-required:"true" yaml:"name" env:"APP_NAME"`
}
