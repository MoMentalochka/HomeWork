package session

import "github.com/MoMentalochka/HomeWork/platform/pkg/cache"

const (
	cacheKeyPrefix = "auth:session:"
)

type repository struct {
	cache cache.RedisClient
}

func NewRepository(cache cache.RedisClient) *repository {
	return &repository{
		cache: cache,
	}
}
