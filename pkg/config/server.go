package config

type Server struct {
	Pprof Pprof `yaml:"pprof" json:"pprof"`
	Grpc  Grpc  `yaml:"grpc" json:"grpc"`
}

type Grpc struct {
	Host       string `yaml:"host" json:"host,omitempty" env:"GRPC_HOST" env-default:"0.0.0.0"`
	Port       int    `yaml:"port" json:"port,omitempty" env:"GRPC_PORT" env-default:"5000"`
	LogPayload *bool  `yaml:"log_payload" json:"log_payload,omitempty" env:"GRPC_LOG_PAYLOAD" env-default:"true"`
}

type Pprof struct {
	Enable *bool  `yaml:"enable" json:"enable,omitempty" env:"PPROF_ENABLE" env-default:"true"`
	Host   string `yaml:"host" json:"host,omitempty" env:"PPROF_HOST" env-default:"0.0.0.0"`
	Port   int    `yaml:"port" json:"port,omitempty" env:"PPROF_PORT" env-default:"6060"`
}
