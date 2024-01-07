package config

type App struct {
	Name string `env-required:"true" yaml:"name" json:"name,omitempty" env:"APP_NAME"`
}
