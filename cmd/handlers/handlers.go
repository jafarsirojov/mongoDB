package handlers

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"mongoDB/internal/record"
	"mongoDB/internal/responses"
	"mongoDB/internal/structs"
	"mongoDB/pkg/reply"
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

	var response structs.Response
	defer reply.Json(w, http.StatusOK, &response)

	list, err := h.recordsService.GetAll(r.Context())
	if err != nil {
		if err == structs.ErrNotFound {
			h.logger.Info("cmd.handlers.GetAll recordsService.GetAll: not found")
			response = responses.NotFound
			return
		}
		h.logger.Error("cmd.handlers.GetAll recordsService.GetAll", zap.Error(err))
		response = responses.InternalErr
		return
	}

	response = responses.Success
	response.Payload = list
}
