package worker

import (
	"context"
	"time"
	"fmt"

	"github.com/seeques/notification-service/internal/queue"
)

type LogProcessor struct {}

type Processor interface {
	Process(ctx context.Context, job *queue.Job)
}

func (lp *LogProcessor) Process(ctx context.Context, job *queue.Job) {
	fmt.Printf("Started working on job id: %s\n", job.ID)

	time.Sleep(time.Second)
}