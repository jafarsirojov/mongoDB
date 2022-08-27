package pkg

import (
	"github.com/jafarsirojov/mongoDB/pkg/cache"
	"github.com/jafarsirojov/mongoDB/pkg/config"
	"github.com/jafarsirojov/mongoDB/pkg/logger"
	"github.com/jafarsirojov/mongoDB/pkg/mongoDB"
	"go.uber.org/fx"
)

var Module = fx.Options(
	config.Module,
	mongoDB.Module,
	logger.Module,
	cache.Module,
)
