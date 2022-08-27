package main

import (
	"github.com/jafarsirojov/mongoDB/cmd/handlers"
	"github.com/jafarsirojov/mongoDB/cmd/job"
	"github.com/jafarsirojov/mongoDB/cmd/router"
	"github.com/jafarsirojov/mongoDB/internal"
	"github.com/jafarsirojov/mongoDB/pkg"
	"go.uber.org/fx"
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
