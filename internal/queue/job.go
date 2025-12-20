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

func (rc *RedisClient) PushJob(job Job) {
	jobJSON, _ := json.Marshal(job)
	rc.RedisClient.LPush(context.Background(), "notifications:pending", jobJSON)
}

func (rc *RedisClient) PopJob(job Job) error {
	result, err := rc.RedisClient.BRPop(context.Background(), 0, "notifications:pending").Result()
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(result[1]), &job)
	return nil
}