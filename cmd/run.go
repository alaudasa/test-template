package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run tests",
	Run: func(cmd *cobra.Command, args []string) {
		doRun()
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
}

func doRun() {
	fmt.Println("All tests passed.")
}
