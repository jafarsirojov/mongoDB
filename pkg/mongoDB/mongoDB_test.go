package mongoDB

import (
	"context"
	"github.com/jafarsirojov/mongoDB/internal/structs"
	"github.com/jafarsirojov/mongoDB/pkg/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func Test_mongoDB_Add(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		config *config.Config
	}
	type args struct {
		ctx    context.Context
		record structs.Record
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"ok",
			fields{
				logger: zap.NewNop(),
				config: config.ProvideConfig()},
			args{
				ctx:    context.Background(),
				record: testRecord[0],
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewDB(Params{
				Logger: tt.fields.logger,
				Config: tt.fields.config,
			})
			if err := m.Add(tt.args.ctx, tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mongoDB_GetAll(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		config *config.Config
	}
	type args struct {
		ctx    context.Context
		filter interface{}
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantRecords []structs.Record
		wantErr     bool
		Err         error
	}{
		{
			"not found",
			fields{
				logger: zap.NewNop(),
				config: config.ProvideConfig()},
			args{
				ctx: context.Background(),
				filter: primitive.D{primitive.E{
					Key:   "name",
					Value: "ueidnwocjp",
				}},
			},
			nil,
			true,
			structs.ErrNotFound,
		}, {
			"one",
			fields{
				logger: zap.NewNop(),
				config: config.ProvideConfig()},
			args{
				ctx: context.Background(),
				filter: primitive.D{primitive.E{
					Key:   "name",
					Value: "name1",
				}},
			},
			testRecord,
			false,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewDB(Params{
				Logger: tt.fields.logger,
				Config: tt.fields.config,
			})
			gotRecords, err := m.GetAll(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr || err != tt.Err {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecords, tt.wantRecords) {
				t.Errorf("GetAll() gotRecords = %v, want %v", gotRecords, tt.wantRecords)
			}
		})
	}
}

func Test_mongoDB_Delete(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		config *config.Config
	}
	type args struct {
		ctx    context.Context
		filter interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"ok",
			fields{
				logger: zap.NewNop(),
				config: config.ProvideConfig()},
			args{
				ctx: context.Background(),
				filter: primitive.D{primitive.E{
					Key:   "name",
					Value: testRecord[0].Name,
				}},
			},
			false,
		},
		{
			"value is nil",
			fields{
				logger: zap.NewNop(),
				config: config.ProvideConfig()},
			args{
				ctx:    context.Background(),
				filter: nil,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewDB(Params{
				Logger: tt.fields.logger,
				Config: tt.fields.config,
			})
			if err := m.Delete(tt.args.ctx, tt.args.filter); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

var testRecord = []structs.Record{
	{
		Name:   "name1",
		Text:   "gyedcenwc3456uwuwcnw",
		Status: 0,
	},
}
