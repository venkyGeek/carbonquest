package cmd

import (
	"fmt"
	"carbonquest/utils"
	"github.com/spf13/cobra"
)

var env, version string
var clean bool

var deployCmd = &cobra.Command{
	Use:   "deploy <function>",
	Short: "Deploy a cloud function to GCP",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		function := args[0]
		if clean {
			fmt.Println("Cleaning and rebuilding...")
			// Example: run go build or clean steps
		}
		utils.DeployFunction(function, env, version)
	},
}

func init() {
	deployCmd.Flags().StringVarP(&env, "env", "e", "dev", "Environment: sandbox, dev, prod")
	deployCmd.Flags().StringVarP(&version, "version", "v", "latest", "Function version")
	deployCmd.Flags().BoolVarP(&clean, "clean", "c", false, "Clean and rebuild before deploy")
}
