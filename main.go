package main

import (
	"time"
	"syscall"
	"os"
	"os/signal"
	"context"

	"github.com/rs/zerolog/log"
	"github.com/seeques/notification-service/internal/queue"
	"github.com/seeques/notification-service/internal/config"
	"github.com/seeques/notification-service/internal/worker"
)

func main() {
	cfg := config.LoadConfig()

	q, err := queue.NewQueue(cfg)
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

	for i := 0; i < 10; i++ {
		q.PushJob(&job)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	processor := worker.NewProcessor()
	pool := worker.NewPool(5, q, processor)

	go func() {
		<-sigs
		log.Info().Msg("Shutdown signal received")
		cancel()
	}()

	pool.Start(ctx)
}