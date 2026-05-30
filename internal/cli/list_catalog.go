package cli

import (
	"github.com/Jud4Mo/multi-templates/internal/catalog"
	"github.com/Jud4Mo/multi-templates/internal/prompt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list-catalog",
	Short: "Shows the valid catalogs",
	RunE: func(cmd *cobra.Command, args []string) error {
		cat := catalog.NewCatalog()
		prompt.PrintCatalog(cat)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
