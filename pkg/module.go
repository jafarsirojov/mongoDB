package pkg

import (
	"go.uber.org/fx"
	"mongoDB/pkg/config"
	"mongoDB/pkg/logger"
	"mongoDB/pkg/mongoDB"
)

var Module = fx.Options(
	config.Module,
	mongoDB.Module,
	logger.Module,
)
