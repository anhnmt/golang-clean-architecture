package config

type Postgres struct {
	Host     string `yaml:"host" env:"POSTGRES_HOST" env-default:"localhost"`
	Port     int    `yaml:"port" env:"POSTGRES_PORT" env-default:"5432"`
	User     string `yaml:"user" env:"POSTGRES_USER" env-default:"postgres"`
	Password string `yaml:"password" env:"POSTGRES_PASSWORD" env-default:"postgres"`
	Database string `yaml:"database" env:"POSTGRES_DB" env-default:"postgres"`
	SSLMode  string `yaml:"ssl_mode" env:"POSTGRES_SSL_MODE" env-default:"disable"`

	MaxConns int32 `yaml:"max_conns" env:"POSTGRES_MAX_CONNS" env-default:"10"`
	MinConns int32 `yaml:"min_conns" env:"POSTGRES_MIN_CONNS" env-default:"1"`

	MaxConnIdleTime string `yaml:"max_conn_idle_time" env:"POSTGRES_MAX_CONN_IDLE_TIME" env-default:"5m"`
	MaxConnLifetime string `yaml:"max_conn_lifetime" env:"POSTGRES_MAX_CONN_LIFETIME" env-default:"5m"`
	ConnTimeout     string `yaml:"conn_timeout" env:"POSTGRES_CONN_TIMEOUT" env-default:"15s"`
}
