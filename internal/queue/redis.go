package queue

import (
	"github.com/redis/go-redis/v9"
	"github.com/seeques/notification-service/internal/config"
)

type Queue struct {
	RedisClient *redis.Client
}

func NewQueue(cfg config.Config) (*Queue, error) {
	opt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)

	q := &Queue{
		RedisClient: client,
	}

	return q, nil
}