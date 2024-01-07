package config

type Postgres struct {
	Host     string `yaml:"host" json:"host,omitempty" env:"POSTGRES_HOST" env-default:"localhost"`
	Port     int    `yaml:"port" json:"port,omitempty" env:"POSTGRES_PORT" env-default:"5432"`
	User     string `yaml:"user" json:"user,omitempty" env:"POSTGRES_USER" env-default:"postgres"`
	Password string `yaml:"password" json:"password,omitempty" env:"POSTGRES_PASSWORD" env-default:"postgres"`
	Database string `yaml:"database" json:"database,omitempty" env:"POSTGRES_DB" env-default:"postgres"`

	SSLMode string `yaml:"ssl_mode" json:"ssl_mode,omitempty" env:"POSTGRES_SSL_MODE" env-default:"disable"`
	Migrate bool   `yaml:"migrate" json:"migrate,omitempty" env:"POSTGRES_MIGRATE" env-default:"true"`

	MaxConns int32 `yaml:"max_conns" json:"max_conns,omitempty" env:"POSTGRES_MAX_CONNS" env-default:"10"`
	MinConns int32 `yaml:"min_conns" json:"min_conns,omitempty" env:"POSTGRES_MIN_CONNS" env-default:"1"`

	MaxConnIdleTime string `yaml:"max_conn_idle_time" json:"max_conn_idle_time,omitempty" env:"POSTGRES_MAX_CONN_IDLE_TIME" env-default:"5m"`
	MaxConnLifetime string `yaml:"max_conn_lifetime" json:"max_conn_lifetime,omitempty" env:"POSTGRES_MAX_CONN_LIFETIME" env-default:"5m"`
	ConnTimeout     string `yaml:"conn_timeout" json:"conn_timeout,omitempty" env:"POSTGRES_CONN_TIMEOUT" env-default:"15s"`
}
