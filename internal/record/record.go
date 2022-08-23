package record

import (
	"context"
	"go.uber.org/fx"
	"mongoDB/internal/structs"
	"mongoDB/pkg/mongoDB"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In
	MongoDB mongoDB.MongoDB
}

type service struct {
	mongoDB mongoDB.MongoDB
}

func New(params Params) RecordsService {
	return &service{mongoDB: params.MongoDB}
}

type RecordsService interface {
	GetAll(ctx context.Context) (records []structs.Record, err error)
}

func (s *service) GetAll(ctx context.Context) (records []structs.Record, err error) {
	records, err = s.mongoDB.GetAll(ctx, nil)
	if err != nil {
		return nil, err
	}

	return records, nil
}
