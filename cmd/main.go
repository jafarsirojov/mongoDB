package main

import (
	"go.uber.org/fx"
	"mongoDB/cmd/handlers"
	"mongoDB/cmd/job"
	"mongoDB/cmd/router"
	"mongoDB/internal"
	"mongoDB/pkg"
)

func main() {
	fx.New(
		internal.Module,
		pkg.Module,
		handlers.Module,
		router.Module,
		job.Module,
	).Run()
}
