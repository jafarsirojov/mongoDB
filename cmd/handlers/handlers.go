package handlers

import (
	"encoding/json"
	"errors"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"mongoDB/internal/record"
	"net/http"
)

var Module = fx.Provide(NewHandler)

type Params struct {
	fx.In
	Logger         *zap.Logger
	RecordsService record.RecordsService
}

type handler struct {
	logger         *zap.Logger
	recordsService record.RecordsService
}

func NewHandler(params Params) RecordsHandler {
	return &handler{
		logger:         params.Logger,
		recordsService: params.RecordsService,
	}
}

type RecordsHandler interface {
	GetAll(http.ResponseWriter, *http.Request)
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	list, err := h.recordsService.GetAll(r.Context())
	if err != nil {
		if err == errors.New("not found") {
			h.logger.Error("cmd.handlers.GetAll recordsService.GetAll: not found")
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(404)
		}
		h.logger.Error("cmd.handlers.GetAll recordsService.GetAll", zap.Error(err))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)
	}

	respByte, err := json.Marshal(list)
	if err != nil {
		h.logger.Error("cmd.handlers.GetAll json.Marshal", zap.Any("records", list), zap.Error(err))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(respByte)
}
