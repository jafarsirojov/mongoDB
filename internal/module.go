package internal

import (
	"go.uber.org/fx"
	"mongoDB/internal/record"
)

var Module = fx.Options(
	record.Module,
)
