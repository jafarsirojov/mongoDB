package cache

import (
	memoryCache "github.com/patrickmn/go-cache"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"mongoDB/pkg/config"
	"time"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In
	Logger *zap.Logger
	Config *config.Config
}

type cache struct {
	logger *zap.Logger
	config *config.Config
	memory *memoryCache.Cache
}

// This project is small and not scaled horizontally, you can do without redis
type MemoryCache interface {
	Set(key string, data interface{}, d time.Duration)
	Get(key string) (interface{}, bool)
	Replace(k string, x interface{}, d time.Duration) error
}

func New() MemoryCache {
	return &cache{memory: memoryCache.New(memoryCache.NoExpiration, memoryCache.NoExpiration)}
}

func (c *cache) Set(k string, x interface{}, d time.Duration) {
	c.memory.Set(k, x, d)
	return
}

func (c *cache) Get(key string) (interface{}, bool) {
	return c.memory.Get(key)
}

func (c *cache) Replace(k string, x interface{}, d time.Duration) error {
	return c.memory.Replace(k, x, d)
}
