package record

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"mongoDB/internal/job"
	"mongoDB/internal/structs"
	"mongoDB/pkg/cache"
	"mongoDB/pkg/mongoDB"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In
	Logger      *zap.Logger
	MongoDB     mongoDB.MongoDB
	MemoryCache cache.MemoryCache
	JobsService job.JobsService
}

type service struct {
	logger      *zap.Logger
	mongoDB     mongoDB.MongoDB
	memoryCache cache.MemoryCache
	jobsService job.JobsService
}

func New(params Params) RecordsService {
	return &service{
		logger:      params.Logger,
		mongoDB:     params.MongoDB,
		memoryCache: params.MemoryCache,
		jobsService: params.JobsService,
	}
}

type RecordsService interface {
	GetAll(ctx context.Context) (records []structs.Record, err error)
	DeleteByName(ctx context.Context, name string) error
	UpdateByName(ctx context.Context, name string, record structs.Record)  error
}

func (s *service) GetAll(ctx context.Context) (records []structs.Record, err error) {

	v, ok := s.memoryCache.Get("allRecords")
	records, _ = v.([]structs.Record)
	if !ok {
		records, err = s.mongoDB.GetAll(ctx, nil)
		if err != nil {
			if err == structs.ErrNotFound {
				s.logger.Info("internal.record.GetAll s.mongoDB.GetAll: not found")
				return nil, err
			}
			s.logger.Error("internal.record.GetAll s.mongoDB.GetAll", zap.Error(err))
			return nil, err
		}

		s.memoryCache.Set("allRecords", records, -1)
	}

	return records, nil
}

func (s *service) DeleteByName(ctx context.Context, name string) error {

	filter := bson.D{primitive.E{Key: "name", Value: name}}

	err := s.mongoDB.Delete(ctx, filter)
	if err != nil {
		s.logger.Error("internal.record.DeleteByName s.mongoDB.Delete", zap.Error(err))
		return err
	}

	s.jobsService.ResetRecordsCache()

	return nil
}

func (s *service) UpdateByName(ctx context.Context, name string, record structs.Record) error {

	filter := bson.D{primitive.E{Key: "name", Value: name}}

	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "status", Value: record.Status},
	}}}

	err := s.mongoDB.Update(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			s.logger.Info("internal.record.Update s.mongoDB.Update", zap.Error(err))
			return err
		}
		s.logger.Error("internal.record.Update s.mongoDB.Update", zap.Error(err))
		return err
	}

	s.jobsService.ResetRecordsCache()

	return nil
}
