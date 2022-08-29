package config

import (
	"encoding/json"
	"fmt"
	"go.uber.org/fx"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	MongoDbUrl                string
	Port                      string
	Version                   string
	ResetRecordsCacheDuration time.Duration
}

var Module = fx.Options(fx.Provide(ProvideConfig))

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	if s, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("bad value %q for %s: %v", s, key, err)
		}
		return time.Duration(v)
	}
	return fallback
}

func (c *Config) Dump(w io.Writer) error {
	fmt.Fprint(w, "config: ")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	return enc.Encode(c)
}

func ProvideConfig() *Config {
	return &Config{
		MongoDbUrl:                getEnv("MONGO_DB_URL", "mongodb://localhost:27017/"),
		Port:                      getEnv("PORT", ":7777"),
		Version:                   getEnv("VERSION", "v1"),
		ResetRecordsCacheDuration: getEnvDuration("RESET_RECORDS_CACHE_DURATION", time.Second),
	}
}
