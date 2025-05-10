package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gcptool",
	Short: "CLI to deploy and describe GCP functions",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}

func init() {
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(describeCmd)
}
