package handlers

import (
	"context"
	"encoding/json"
	"go.uber.org/fx"
	"log"
	"mongoDB/internal/record"
	"net/http"
)

var Module = fx.Provide(NewHandler)

type Params struct {
	fx.In
	RecordsService record.RecordsService
}

type handler struct {
	recordsService record.RecordsService
}

func NewHandler(params Params) RecordsHandler {
	return &handler{
		recordsService: params.RecordsService,
	}
}

type RecordsHandler interface {
	GetAll(http.ResponseWriter, *http.Request)
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	list, err := h.recordsService.GetAll(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	respByte, err := json.Marshal(list)
	if err != nil {
		log.Fatal("json.Marshal(list):", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(respByte)
}
