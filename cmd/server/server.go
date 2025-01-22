package server

import (
	"fmt"
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
)

func serverRun(cmd *cobra.Command, args []string) error {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	server := &http.Server{
		Addr:              address,
		Handler:           nil,
		ReadHeaderTimeout: 1 * time.Minute,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	ServerCmd.RunE = serverRun
	ServerCmd.PersistentFlags().StringVarP(&address, "server", "s", ":4000", "Server address (default \":4000\")")
}
