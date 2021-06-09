package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "v0.0.0+unknown"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print ikscc version",
	Long:  "Print the version of the ikscc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("ikscc version: %s", Version)
	},
}
