package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the api-kit version",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("api-kit version", Version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
