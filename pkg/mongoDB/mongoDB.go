package mongoDB

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/fx"
	"log"
	"mongoDB/internal/structs"
	"mongoDB/pkg/config"
)

var Module = fx.Provide(NewDB)

type Params struct {
	Config *config.Config
}

type mongoDB struct {
	config *config.Config
}

func NewDB(params Params) MongoDB {
	initClient()
	return &mongoDB{config: params.Config}
}

type MongoDB interface {
	Add(ctx context.Context, record structs.Record) error
	GetAll(ctx context.Context, filter interface{}) (records []structs.Record, err error)
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
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		var record structs.Record
		err = cur.Decode(&record)
		if err != nil {
			return nil, err
		}

		records = append(records, record)
	}

	if len(records) == 0 {
		return nil, errors.New("not found")
	}

	return records, nil
}
