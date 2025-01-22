package cmd

import (
	"github.com/arahmandanu/sinau_go_craft/cmd/server"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "webhook-receiver",
	Short: "Qontak Chat Webhook",
	Long:  "Qontak chat webhook - receive and process incoming message from channel integration.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_ = rootCmd.Help()
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(server.ServerCmd)
}
