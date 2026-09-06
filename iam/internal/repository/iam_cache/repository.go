package iam_cache

import "github.com/MoMentalochka/HomeWork/platform/pkg/cache"

const (
	cacheKeyPrefix = "ufo:sighting:"
)

type repository struct {
	cache cache.RedisClient
}

func NewRepository(cache cache.RedisClient) *repository {
	return &repository{
		cache: cache,
	}
}
