package cmd

import (
	"github.com/hugogarcia/go-kakfa/api"
	"github.com/spf13/cobra"
)

var apiCommand = &cobra.Command{
	Use:   "api",
	Short: "API with producer",
	RunE:  startServer,
}

func init() {
	rootCmd.AddCommand(apiCommand)
}

func startServer(cmd *cobra.Command, args []string) error {
	return api.Execute()
}
