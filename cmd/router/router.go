package router

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"mongoDB/cmd/handlers"
	"mongoDB/pkg/config"
	"net/http"
)

var Module = fx.Invoke(NewRouter)

type Params struct {
	fx.In
	Lifecycle      fx.Lifecycle
	Config         *config.Config
	RecordsHandler handlers.RecordsHandler
}

func NewRouter(params Params) {
	router := mux.NewRouter()

	router.HandleFunc("/all", params.RecordsHandler.GetAll).Methods("GET")
	router.HandleFunc("/delete/{name}", params.RecordsHandler.DeleteByName).Methods("DELETE")

	server := http.Server{
		Addr:    params.Config.Port,
		Handler: router,
	}

	params.Lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				fmt.Println("Application started")
				go server.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				fmt.Println("Application stopped")
				return server.Shutdown(ctx)
			},
		},
	)
}
