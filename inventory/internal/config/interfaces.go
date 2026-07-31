package config

type InventoryGRPCConfig interface {
	Address() string
}

type MongoConfig interface {
	URI() string
	DB_Name() string
}

type LoggerConfig interface {
	Level() string
	AsJson() bool
}
