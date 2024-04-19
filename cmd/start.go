package cmd

import (
	"github.com/oneness/erc-4337-api/internal/start"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a JSON-RPC server",
	Long:  "The start command will run a JSON-RPC server to enable processing of UserOperations api.",
	Run: func(cmd *cobra.Command, args []string) {
		start.Server()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
