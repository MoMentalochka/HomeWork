package config

type InventoryGRPCConfig interface {
	Address() string
}

type MongoConfig interface {
	URI() string
	DB_Name() string
}
