package record

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"mongoDB/internal/structs"
	"mongoDB/pkg/mongoDB"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In
	Logger  *zap.Logger
	MongoDB mongoDB.MongoDB
}

type service struct {
	logger  *zap.Logger
	mongoDB mongoDB.MongoDB
}

func New(params Params) RecordsService {
	return &service{
		logger:  params.Logger,
		mongoDB: params.MongoDB,
	}
}

type RecordsService interface {
	GetAll(ctx context.Context) (records []structs.Record, err error)
}

func (s *service) GetAll(ctx context.Context) (records []structs.Record, err error) {
	records, err = s.mongoDB.GetAll(ctx, nil)
	if err != nil {
		if err == structs.ErrNotFound {
			s.logger.Info("internal.record.GetAll s.mongoDB.GetAll: not found")
			return nil, err
		}
		s.logger.Error("internal.record.GetAll s.mongoDB.GetAll", zap.Error(err))
		return nil, err
	}

	return records, nil
}
