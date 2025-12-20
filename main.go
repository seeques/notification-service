package main

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/seeques/notification-service/internal/queue"
	"github.com/seeques/notification-service/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	rc, err := queue.NewRedisClient(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create new redis client")
	}

	ctx := context.Background()
	
	err = rc.RedisClient.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to set a string")
	}

	val, err := rc.RedisClient.Get(ctx, "foo").Result()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get result")
	}
	log.Info().Str("foo:", val).Msg("")
}