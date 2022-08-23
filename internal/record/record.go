package record

import (
	"context"
	"go.uber.org/fx"
	"mongoDB/internal/structs"
	"mongoDB/pkg/mongoDB"
)

var Module = fx.Provide(New)

type Params struct {
	MongoDB mongoDB.MongoDB
}

type service struct {
	mongoDB mongoDB.MongoDB
}

func New(params Params) Record {
	return &service{mongoDB: params.MongoDB}
}

type Record interface {
	GetAll(ctx context.Context) (records []structs.Record, err error)
}

func (s *service) GetAll(ctx context.Context) (records []structs.Record, err error) {
	records, err = s.mongoDB.GetAll(ctx, nil)
	if err != nil {
		return nil, err
	}

	return records, nil
}
