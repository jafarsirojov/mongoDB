package mongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"mongoDB/internal/structs"
	"mongoDB/pkg/config"
)

var Module = fx.Provide(NewDB)

type Params struct {
	fx.In
	Logger *zap.Logger
	Config *config.Config
}

type mongoDB struct {
	logger *zap.Logger
	config *config.Config
}

func NewDB(params Params) MongoDB {
	initClient(params)
	return &mongoDB{
		logger: params.Logger,
		config: params.Config,
	}
}

type MongoDB interface {
	Add(ctx context.Context, record structs.Record) error
	GetAll(ctx context.Context, filter interface{}) (records []structs.Record, err error)
	Delete(ctx context.Context, filter interface{}) error
}

func (m *mongoDB) Add(ctx context.Context, record structs.Record) error {
	_, err := collection.InsertOne(ctx, record)
	return err
}

func (m *mongoDB) GetAll(ctx context.Context, filter interface{}) (records []structs.Record, err error) {

	if filter == nil {
		filter = bson.D{{}} // без филтра
	}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		m.logger.Error("pkg.mongoDB.GetAll collection.Find", zap.Error(err))
		return nil, err
	}

	for cur.Next(ctx) {
		var record structs.Record
		err = cur.Decode(&record)
		if err != nil {
			m.logger.Error("pkg.mongoDB.GetAll cur.Decode", zap.Any("cur", cur), zap.Error(err))
			return nil, err
		}

		records = append(records, record)
	}

	if len(records) == 0 {
		return nil, structs.ErrNotFound
	}

	return records, nil
}
func (m *mongoDB) Delete(ctx context.Context, filter interface{}) error {
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		m.logger.Error("pkg.mongoDB.GetAll cur.Decode", zap.Any("filter", filter), zap.Error(err))
		return err
	}
	return nil
}
