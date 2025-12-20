package worker

import (
	"context"
	"sync"

	"github.com/seeques/notification-service/internal/queue"
	"github.com/rs/zerolog/log"
)

type Pool struct {
	workers int
	queue *queue.Queue
	processor Processor
}

func NewPool(wrs int, q *queue.Queue, proc Processor) *Pool {
	return &Pool{
		workers: wrs,
		queue: q,
		processor: proc,
	}
}

func (p *Pool) Start() {
	var wg sync.WaitGroup

	for i := 0; i < p.workers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			Work(p.queue, id)
		}(i)
	}

	wg.Wait()
}

func Work(q *queue.Queue, id int) {
	var lp LogProcessor
	ctx := context.Background()
	for {
		log.Info().Interface("Worker:", id).Msg("Picked up job")
		retJob, err := q.PopJob()
		if err != nil || retJob == nil {
			break
		}

		lp.Process(ctx, retJob)
		log.Info().Interface("Worker:", id).Msg("Finished the job")
	}
}