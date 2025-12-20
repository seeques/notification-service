package main

import (
	"time"

	"github.com/rs/zerolog/log"
	"github.com/seeques/notification-service/internal/queue"
	"github.com/seeques/notification-service/internal/config"
	"github.com/seeques/notification-service/internal/worker"
)

func main() {
	cfg := config.LoadConfig()

	rc, err := queue.NewRedisClient(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create new redis client")
	}
	
	job := queue.Job{
		ID: "0",
		Channel: "email@email.io",
		Recipient: "recipient@email.io",
		Subject: "test",
		Body: "test body",
		Attempts: 0,
		CreatedAt: time.Now(),
	}

	for i := 0; i < 3; i++ {
		rc.PushJob(&job)
	}

	log.Info().Interface("Job:", job).Msg("Job pushed")

	worker.Work(rc)
}