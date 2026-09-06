package iam

import (
	"time"

	"github.com/MoMentalochka/HomeWork/iam/internal/repository"
)

type iamService struct {
	repository      repository.IamRepository
	cacheRepository repository.IamCacheRepository
	cacheTTL        time.Duration
}

func NewIamService(rep repository.IamRepository, cache repository.IamCacheRepository, cacheTTL time.Duration) *iamService {
	return &iamService{
		rep,
		cache,
		cacheTTL,
	}
}
