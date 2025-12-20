package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisURL string
}

func LoadConfig() Config {
	godotenv.Load()

	return Config{
		RedisURL: os.Getenv("REDIS_URL"),
	}
}