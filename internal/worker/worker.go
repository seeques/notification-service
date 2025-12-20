package worker

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/seeques/notification-service/internal/queue"
)

func Work(rc *queue.RedisClient) {
	var lp LogProcessor
	ctx := context.Background()
	for {
		retJob, err := rc.PopJob()
		if err != nil {
			log.Error().Err(err).Msg("Error processing job")
		}

		lp.Process(ctx, retJob)
	}
}