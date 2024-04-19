package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "erc4337-api-server",
	Short: "ERC-4337 API Server",
	Long:  "A JSON-RPC server for ERC-4337 api server.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
