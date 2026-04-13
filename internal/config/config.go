package config

import "os"

type Config struct {
	DbConnection string
	CacheHost    string
	CachePort    string
}

func Init() *Config {
	var cfg Config
	cfg.DbConnection = os.Getenv("DB_CONNECTION")
	cfg.CacheHost = os.Getenv("CACHE_HOST")
	cfg.CachePort = os.Getenv("CACHE_PORT")

	return &cfg
}
