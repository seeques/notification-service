package worker

import (
	"context"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/seeques/notification-service/internal/queue"
)

type Pool struct {
	workers   int
	queue     *queue.Queue
	processor Processor
}

func NewPool(wrs int, q *queue.Queue, proc Processor) *Pool {
	return &Pool{
		workers:   wrs,
		queue:     q,
		processor: proc,
	}
}

func (p *Pool) Start(ctx context.Context) {
	var wg sync.WaitGroup

	for i := 0; i < p.workers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			p.work(ctx, id)
		}(i)
	}

	wg.Wait()
}

func (p *Pool) work(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			log.Info().Interface("Worker:", id).Msg("Picked up job")
			retJob, err := p.queue.PopJob()
			if err != nil || retJob == nil {
				continue
			}

			p.processor.Process(ctx, retJob)
			log.Info().Interface("Worker:", id).Msg("Finished the job")
		}
	}
}
