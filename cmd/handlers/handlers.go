package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jafarsirojov/mongoDB/internal/record"
	"github.com/jafarsirojov/mongoDB/internal/responses"
	"github.com/jafarsirojov/mongoDB/internal/structs"
	"github.com/jafarsirojov/mongoDB/pkg/reply"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/fx"
	"go.uber.org/zap"
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
	DeleteByID(http.ResponseWriter, *http.Request)
	UpdateByID(http.ResponseWriter, *http.Request)
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

func (h *handler) DeleteByID(w http.ResponseWriter, r *http.Request) {

	var response structs.Response
	defer reply.Json(w, http.StatusOK, &response)

	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		h.logger.Error("cmd.handlers.DeleteByName primitive.ObjectIDFromHex: check 'id' params",
			zap.String("id", mux.Vars(r)["id"]), zap.Error(err))
		response = responses.BadRequest
		return
	}

	err = h.recordsService.DeleteByID(r.Context(), id)
	if err != nil {
		if err == structs.ErrNotFound {
			h.logger.Info("cmd.handlers.DeleteByID recordsService.DeleteByID: not found")
			response = responses.NotFound
			return
		}
		h.logger.Error("cmd.handlers.DeleteByID recordsService.DeleteByID", zap.Error(err))
		response = responses.InternalErr
		return
	}

	response = responses.Success
}

func (h *handler) UpdateByID(w http.ResponseWriter, r *http.Request) {

	var (
		response structs.Response
		request  structs.Record
	)

	defer reply.Json(w, http.StatusOK, &response)

	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		h.logger.Error("cmd.handlers.UpdateByID primitive.ObjectIDFromHex: check 'id' params",
			zap.String("id", mux.Vars(r)["id"]), zap.Error(err))
		response = responses.BadRequest
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.logger.Error("cmd.handlers.UpdateByID json.NewDecoder(r.Body).Decode(&request)",
			zap.Any("request", request), zap.Any("r.Body", r.Body), zap.Error(err))
		response = responses.BadRequest
		return
	}

	err = h.recordsService.UpdateByID(r.Context(), id, request)
	if err != nil {
		if err == structs.ErrNotFound {
			h.logger.Info("cmd.handlers.UpdateByID recordsService.UpdateByID: not found")
			response = responses.NotFound
			return
		}
		h.logger.Error("cmd.handlers.UpdateByID recordsService.UpdateByID", zap.Error(err))
		response = responses.InternalErr
		return
	}

	response = responses.Success
}
