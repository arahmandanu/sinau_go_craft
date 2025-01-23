package worker

import (
	"errors"
	"fmt"
	"github.com/arahmandanu/sinau_go_craft/config"
	"github.com/gocraft/work"
	redis2 "github.com/gomodule/redigo/redis"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var (
	address   string
	ServerCmd = &cobra.Command{
		Use:   "worker",
		Short: "Run chat webhook worker server",
		Long:  "Run chat webhook worker server to process background job.",
	}
	rdb       *redis.Client
	redisPool = &redis2.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis2.Conn, error) {
			return redis2.Dial("tcp", fmt.Sprintf(":%d", config.GetRedisPort()))
		},
	}
)

type Context struct {
}

func preServerRun(cmd *cobra.Command, args []string) error {
	config.Init()
	rOpt, err := config.InitRedis()
	if err != nil {
		return err
	}

	rdb, err = config.CallRedis(rOpt)
	return nil
}

func serverRun(cmd *cobra.Command, args []string) error {
	pool := work.NewWorkerPool(Context{}, 10, "my_app_namespace", redisPool)
	pool.Middleware((*Context).Log)

	pool.Job("adrian_job", TestDrive)

	//pool.JobWithOptions("adrian_job", work.JobOptions{Priority: 10, MaxFails: 1}, (*Context).Export)
	pool.JobWithOptions("adrian_job", work.JobOptions{Priority: 10, MaxFails: 10}, (*Context).Export)

	pool.Start()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	defer func() {
		println("Shutting down worker...")
		pool.Stop()
		println("successfully shut down worker...")
	}()

	return nil
}

func TestDrive(job *work.Job) error {
	addr := job.ArgString("address")
	subject := job.ArgString("subject")
	if err := job.ArgError(); err != nil {
		return err
	}
	fmt.Println("testing" + addr + subject)
	fmt.Println("Testing drive")
	return nil
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting job: ", job.Name)
	return errors.New("error")
	//return next()
}

func (c *Context) Export(job *work.Job) error {
	return nil
}

func init() {
	ServerCmd.PreRunE = preServerRun
	ServerCmd.RunE = serverRun
	ServerCmd.PersistentFlags().StringVarP(&address, "server", "s", ":4001", "Server address (default \":4001\")")
}
