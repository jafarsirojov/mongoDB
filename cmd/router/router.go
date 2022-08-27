package router

import (
	"context"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"mongoDB/cmd/handlers"
	"mongoDB/pkg/config"
	"net/http"
)

var Module = fx.Invoke(NewRouter)

type Params struct {
	fx.In
	Lifecycle      fx.Lifecycle
	Config         *config.Config
	Logger         *zap.Logger
	RecordsHandler handlers.RecordsHandler
}

func NewRouter(params Params) {
	router := mux.NewRouter()

	version := params.Config.Version
	baseUrl := "/api/record/" + version

	router.HandleFunc(baseUrl+"/all", params.RecordsHandler.GetAll).Methods("GET")
	router.HandleFunc(baseUrl+"/delete/{id}", params.RecordsHandler.DeleteByID).Methods("DELETE")
	router.HandleFunc(baseUrl+"/update/{id}", params.RecordsHandler.UpdateByID).Methods("PUT")

	server := http.Server{
		Addr:    params.Config.Port,
		Handler: router,
	}

	params.Lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				params.Logger.Debug("Application started")
				go server.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				params.Logger.Debug("Application stopped")
				return server.Shutdown(ctx)
			},
		},
	)
}
