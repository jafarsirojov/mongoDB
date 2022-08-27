package internal

import (
	"github.com/jafarsirojov/mongoDB/internal/job"
	"github.com/jafarsirojov/mongoDB/internal/record"
	"go.uber.org/fx"
)

var Module = fx.Options(
	record.Module,
	job.Module,
)
