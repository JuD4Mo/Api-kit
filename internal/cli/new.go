package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	newLang      string
	newFramework string
	newArch      string
	newName      string
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Generate a project from a template",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("New Project:")
		fmt.Println(" lang:      ", newLang)
		fmt.Println(" framework: ", newFramework)
		fmt.Println(" arch:      ", newArch)
		fmt.Println(" name:      ", newName)
	},
}

// function init executes automatically when loading the package
func init() {
	newCmd.Flags().StringVar(&newLang, "lang", "", "Language (go, java, ts...)")
	newCmd.Flags().StringVar(&newFramework, "framework", "", "Framework (gin, echo, spring, express...)")
	newCmd.Flags().StringVar(&newArch, "arch", "", "Architecture (layered, hexagonal, clean...)")
	newCmd.Flags().StringVar(&newName, "name", "", "Name of the project to be created")

	rootCmd.AddCommand(newCmd)
}
