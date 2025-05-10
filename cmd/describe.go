package cmd

import (
	"carbonquest/utils"
	"github.com/spf13/cobra"
)


var describeCmd = &cobra.Command{
	Use:   "describe <function>",
	Short: "Describe a deployed cloud function",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils.DescribeFunction(args[0])
	},
}
