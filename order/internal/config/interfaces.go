package config

import "time"

type OrderHttpConfig interface {
	Address() string
	Port() string
	ReadTimeOut() time.Duration
}

type PostgresConfig interface {
	URI() string
	MigrationsDir() string
}

type PaymentConfig interface {
	Address() string
}

type InventoryConfig interface {
	Address() string
}

type LoggerConfig interface {
	Level() string
	AsJson() bool
}
