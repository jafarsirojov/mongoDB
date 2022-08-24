package handlers

import (
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"mongoDB/internal/record"
	"mongoDB/internal/responses"
	"mongoDB/internal/structs"
	"mongoDB/pkg/reply"
	"net/http"
	"strings"
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
	DeleteByName(http.ResponseWriter, *http.Request)
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

func (h *handler) DeleteByName(w http.ResponseWriter, r *http.Request) {

	var response structs.Response
	defer reply.Json(w, http.StatusOK, &response)

	name := mux.Vars(r)["name"]
	if len(strings.TrimSpace(name)) == 0 {
		h.logger.Error("cmd.handlers.DeleteByName check 'name' params", zap.String("name", name))
		response = responses.BadRequest
		return
	}

	err := h.recordsService.DeleteByName(r.Context(), name)
	if err != nil {
		if err == structs.ErrNotFound {
			h.logger.Info("cmd.handlers.DeleteByID recordsService.GetAll: not found")
			response = responses.NotFound
			return
		}
		h.logger.Error("cmd.handlers.DeleteByID recordsService.GetAll", zap.Error(err))
		response = responses.InternalErr
		return
	}

	response = responses.Success
}
