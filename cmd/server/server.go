package server

import (
	"fmt"
	"github.com/arahmandanu/sinau_go_craft/config"
	"github.com/arahmandanu/sinau_go_craft/pkg/background_job"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

var (
	address   string
	ServerCmd = &cobra.Command{
		Use:   "server",
		Short: "Run chat webhook server",
		Long:  "Run chat webhook server to process incoming message.",
	}
	rdb *redis.Client
)

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
	var err error
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		job, err := background_job.InitJob("adrian_job", map[string]interface{}{"adrian": 01})
		if err != nil {
			fmt.Println("error", err)
		} else {
			fmt.Println(job.ID)
		}
	})

	server := &http.Server{
		Addr:              address,
		Handler:           nil,
		ReadHeaderTimeout: 1 * time.Minute,
	}

	err = server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	ServerCmd.PreRunE = preServerRun
	ServerCmd.RunE = serverRun
	ServerCmd.PersistentFlags().StringVarP(&address, "server", "s", ":4000", "Server address (default \":4000\")")
}
