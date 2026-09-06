package app

import (
	"context"
	"database/sql"

	iamV1API "github.com/MoMentalochka/HomeWork/iam/internal/api/iam/v1"
	"github.com/MoMentalochka/HomeWork/iam/internal/config"
	"github.com/MoMentalochka/HomeWork/iam/internal/migrator"
	"github.com/MoMentalochka/HomeWork/iam/internal/repository"
	iamRepository "github.com/MoMentalochka/HomeWork/iam/internal/repository/iam"
	iamCacheRepository "github.com/MoMentalochka/HomeWork/iam/internal/repository/iam_cache"
	"github.com/MoMentalochka/HomeWork/iam/internal/service"
	iamService "github.com/MoMentalochka/HomeWork/iam/internal/service/iam"
	"github.com/MoMentalochka/HomeWork/platform/pkg/cache"
	"github.com/MoMentalochka/HomeWork/platform/pkg/cache/redis"
	"github.com/MoMentalochka/HomeWork/platform/pkg/closer"
	"github.com/MoMentalochka/HomeWork/platform/pkg/logger"
	iamV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/iam/v1"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

type diContainer struct {
	iamAPI       iamV1.IamServiceServer
	iamService   service.IamService
	iamRepo      repository.IamRepository
	iamCacheRepo repository.IamCacheRepository

	postgresDB *sql.DB

	redisPool   *redigo.Pool
	redisClient cache.RedisClient
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) IamV1API(ctx context.Context) iamV1.IamServiceServer {
	if d.iamAPI == nil {
		d.iamAPI = iamV1API.NewApi(d.Service(ctx))
	}
	return d.iamAPI
}

func (d *diContainer) Service(ctx context.Context) service.IamService {
	if d.iamService == nil {
		d.iamService = iamService.NewIamService(
			d.Repository(ctx),
			d.CacheRepository(),
			config.AppConfig().Redis.CacheTTL(),
		)
	}

	return d.iamService
}

func (d *diContainer) Repository(ctx context.Context) repository.IamRepository {
	if d.iamRepo == nil {
		d.iamRepo = iamRepository.NewIamRepository(d.PostgresDB(ctx))
	}

	return d.iamRepo
}

func (d *diContainer) CacheRepository() repository.IamCacheRepository {
	if d.iamCacheRepo == nil {
		d.iamCacheRepo = iamCacheRepository.NewRepository(d.RedisClient())
	}

	return d.iamCacheRepo
}

func (d *diContainer) RedisPool() *redigo.Pool {
	if d.redisPool == nil {
		d.redisPool = &redigo.Pool{
			MaxIdle:     config.AppConfig().Redis.MaxIdle(),
			IdleTimeout: config.AppConfig().Redis.IdleTimeout(),
			DialContext: func(ctx context.Context) (redigo.Conn, error) {
				return redigo.DialContext(ctx, "tcp", config.AppConfig().Redis.Address())
			},
		}
	}

	return d.redisPool
}

func (d *diContainer) RedisClient() cache.RedisClient {
	if d.redisClient == nil {
		d.redisClient = redis.NewClient(d.RedisPool(), logger.Logger(), config.AppConfig().Redis.ConnectionTimeout())
	}

	return d.redisClient
}

func (d *diContainer) PostgresDB(ctx context.Context) *sql.DB {
	if d.postgresDB == nil {
		// Создаем соединение с базой данных
		con, err := pgx.Connect(ctx, config.AppConfig().Postgres.URI())
		if err != nil {
			logger.Error(ctx, "failed to connect postgres: %v\n", zap.Error(err))
			return nil
		}
		// Проверяем, что соединение с базой установлено
		err = con.Ping(ctx)
		if err != nil {
			logger.Error(ctx, "failed to ping postgres: %v\n", zap.Error(err))
			return nil
		}

		closer.AddNamed("PostgresDB client", func(ctx context.Context) error {
			return con.Close(ctx)
		})

		d.postgresDB = stdlib.OpenDB(*con.Config().Copy())
		// Инициализируем мигратор
		migrationRunner := migrator.NewMigrator(stdlib.OpenDB(*con.Config().Copy()), config.AppConfig().Postgres.MigrationsDir())
		// Раскатываем миграции
		err = migrationRunner.Up()
		if err != nil {
			logger.Error(ctx, "failed to run migrations: %v\n", zap.Error(err))
			return nil
		}

	}

	return d.postgresDB
}
