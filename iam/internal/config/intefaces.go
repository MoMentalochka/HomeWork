package config

import "time"

type IamGRPCConfig interface {
	Address() string
}

type PostgresConfig interface {
	URI() string
	MigrationsDir() string
}

type LoggerConfig interface {
	Level() string
	AsJson() bool
}

type RedisConfig interface {
	Address() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
	CacheTTL() time.Duration
}
