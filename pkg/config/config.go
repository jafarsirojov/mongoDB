package config

import (
	"encoding/json"
	"fmt"
	"go.uber.org/fx"
	"io"
	"os"
)

type Config struct {
	MongoDbUrl string
}

var Module = fx.Options(fx.Provide(ProvideConfig))

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
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
		MongoDbUrl: getEnv("MONGO_DB_URL", "mongodb://localhost:27017/"),
	}
}
