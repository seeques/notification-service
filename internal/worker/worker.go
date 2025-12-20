package worker

import (
	"context"
	"sync"

	"github.com/seeques/notification-service/internal/queue"
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

func (p *Pool) Start(q *queue.Queue) {
	var wg sync.WaitGroup

	for i := 0; i < p.workers; i++ {
		wg.Add(1)
		defer wg.Done()
		go Work(q)
	}

	wg.Wait()
}

func Work(q *queue.Queue) {
	var lp LogProcessor
	ctx := context.Background()
	for {
		retJob, err := q.PopJob()
		if err != nil || retJob == nil {
			break
		}

		lp.Process(ctx, retJob)
	}
}