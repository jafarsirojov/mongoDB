package mongoDB

import (
	"context"
	"github.com/jafarsirojov/mongoDB/internal/structs"
	"github.com/jafarsirojov/mongoDB/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"time"
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
	Update(ctx context.Context, filter, update interface{}) error
}

func (m *mongoDB) Add(ctx context.Context, record structs.Record) error {
	record.ID = primitive.NewObjectID()
	record.CreatedAt = time.Now()
	record.UpdatedAt = time.Now()
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
		m.logger.Error("pkg.mongoDB.Delete collection.DeleteOne",
			zap.Any("filter", filter), zap.Error(err))
		return err
	}
	return nil
}

func (m *mongoDB) Update(ctx context.Context, filter, update interface{}) error {
	err := collection.FindOneAndUpdate(ctx, filter, update).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			m.logger.Info("pkg.mongoDB.Update collection.FindOneAndUpdate not found document",
				zap.Any("filter", filter), zap.Any("update", update))
			return structs.ErrNotFound
		}
		m.logger.Error("pkg.mongoDB.Update collection.FindOneAndUpdate",
			zap.Any("filter", filter), zap.Any("update", update), zap.Error(err))
		return err
	}

	return nil
}
