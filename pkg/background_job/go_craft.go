package background_job

import (
	"fmt"
	"github.com/arahmandanu/sinau_go_craft/config"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"log"
)

// Make a redis pool
var (
	redisPool = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf(":%d", config.GetRedisPort()))
		},
	}
	enqueuer = work.NewEnqueuer("my_app_namespace", redisPool)
)

func InitJob(jobName string, args work.Q) (*work.Job, error) {
	job, err := enqueuer.Enqueue(jobName, args)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return job, nil
}
