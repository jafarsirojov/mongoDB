package job

import (
	"context"
	"github.com/jafarsirojov/mongoDB/internal/structs"
	"github.com/jafarsirojov/mongoDB/pkg/cache"
	"github.com/jafarsirojov/mongoDB/pkg/mongoDB"
	memoryCache "github.com/patrickmn/go-cache"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In
	Logger      *zap.Logger
	MongoDB     mongoDB.MongoDB
	MemoryCache cache.MemoryCache
}

type service struct {
	logger      *zap.Logger
	mongoDB     mongoDB.MongoDB
	memoryCache cache.MemoryCache
}

func New(params Params) JobsService {
	return &service{
		logger:      params.Logger,
		mongoDB:     params.MongoDB,
		memoryCache: params.MemoryCache,
	}
}

type JobsService interface {
	ResetRecordsCache()
}

func (s *service) ResetRecordsCache() {
	records, err := s.mongoDB.GetAll(context.Background(), nil)
	if err != nil {
		if err == structs.ErrBadRequest {
			s.logger.Info("internal.job.ResetRecordsCache mongoDB.GetAll: not found records")
			return
		}
		s.logger.Error("internal.job.ResetRecordsCache mongoDB.GetAll", zap.Error(err))
		return
	}

	s.memoryCache.Set("allRecords", records, memoryCache.NoExpiration)

	s.logger.Debug("internal.job.ResetRecordsCache cache updated", zap.Error(err))

	return
}
