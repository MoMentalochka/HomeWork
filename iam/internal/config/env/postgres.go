package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type postgresEnvConfig struct {
	Host         string `env:"POSTGRES_HOST,required"`
	Port         string `env:"POSTGRES_PORT,required"`
	ExtPort      string `env:"EXTERNAL_POSTGRES_PORT,required"`
	User         string `env:"POSTGRES_USER,required"`
	Password     string `env:"POSTGRES_PASSWORD,required"`
	Database     string `env:"POSTGRES_DB,required"`
	MigrationDir string `env:"MIGRATION_DIRECTORY,required"`
}

type postgresConfig struct {
	cfg postgresEnvConfig
}

func NewPostgresConfig() (*postgresConfig, error) {
	var cfg postgresEnvConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &postgresConfig{cfg: cfg}, nil
}

func (p *postgresConfig) URI() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		p.cfg.User,
		p.cfg.Password,
		p.cfg.Host,
		p.cfg.Port,
		p.cfg.Database,
	)
}

func (p *postgresConfig) MigrationsDir() string {
	return p.cfg.MigrationDir
}
