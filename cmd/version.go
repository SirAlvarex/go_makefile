
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var (
	version string
	buildDate string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of go_makefile",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("go_makefile: %s - %s\n", version, buildDate)
	},
}
