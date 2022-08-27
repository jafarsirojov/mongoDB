package internal

import (
	"go.uber.org/fx"
	"mongoDB/internal/job"
	"mongoDB/internal/record"
)

var Module = fx.Options(
	record.Module,
	job.Module,
)
