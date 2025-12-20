package queue

import (
	"time"
	"encoding/json"
	"context"
)

type Job struct {
	ID        string    `json:"id"`
    Channel   string    `json:"channel"`
    Recipient string    `json:"recipient"`
    Subject   string    `json:"subject"`
    Body      string    `json:"body"`
    Attempts  int       `json:"attempts"`
    CreatedAt time.Time `json:"created_at"`
}

func (q *Queue) PushJob(job *Job) {
	jobJSON, _ := json.Marshal(job)
	q.RedisClient.LPush(context.Background(), "notifications:pending", jobJSON)
}

func (q *Queue) PopJob() (*Job, error) {
	var job *Job
	result, err := q.RedisClient.BRPop(context.Background(), time.Second * 2, "notifications:pending").Result()
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(result[1]), &job)
	return job, nil
}