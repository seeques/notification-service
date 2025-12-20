package queue

import (
	"github.com/redis/go-redis/v9"
	"github.com/seeques/notification-service/internal/config"
)

type RedisClient struct {
	RedisClient *redis.Client
}

func NewRedisClient(cfg config.Config) (*RedisClient, error) {
	opt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)

	rc := &RedisClient{
		RedisClient: client,
	}

	return rc, nil
}