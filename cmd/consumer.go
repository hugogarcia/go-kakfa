package cmd

import (
	"context"

	"github.com/hugogarcia/go-kakfa/kafka"
	"github.com/spf13/cobra"
)

var consummerCommand = &cobra.Command{
	Use:   "consumer",
	Short: "Run kafka consumer",
	RunE:  startConsume,
}

func init() {
	rootCmd.AddCommand(consummerCommand)
}

func startConsume(cmd *cobra.Command, args []string) error {
	return kafka.Consume(context.Background(), kafka.UsersTopic)
}
