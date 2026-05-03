package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mtp",
	Short: "Generates API templates in certain languages and frameworks",
	Long: `multi-templates is a CLI for scaffolding API projects 
by combining language, framework, project type and architecture.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
