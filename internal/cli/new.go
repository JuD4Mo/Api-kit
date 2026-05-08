package cli

import (
	"fmt"

	"github.com/Jud4Mo/multi-templates/internal/catalog"
	"github.com/Jud4Mo/multi-templates/internal/prompt"
	"github.com/spf13/cobra"
)

var (
	newLang        string
	newFramework   string
	newArch        string
	newProjectType string
	newName        string
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Generate a project from a template",
	RunE: func(cmd *cobra.Command, args []string) error {
		input := prompt.NewProjectInputFromFlags(newLang, newFramework, newProjectType, newArch, newName)
		cat := catalog.NewCatalog()

		if input.Language == "" {
			promptInput, err := prompt.AskNewProjectInput(cat)
			if err != nil {
				return err
			}
			input.Language = promptInput.Language
		}

		err := validateProjectInput(cat, input)
		if err != nil {
			return err
		}

		fmt.Println("project configuration is valid")
		fmt.Println(" language:     ", input.Language)
		fmt.Println(" framework:    ", input.Framework)
		fmt.Println(" type:         ", input.ProjectType)
		fmt.Println(" architecture: ", input.Architecture)
		fmt.Println(" name:         ", input.Name)
		return nil
	},
}

// function init executes automatically when loading the package
func init() {
	newCmd.Flags().StringVar(&newLang, "lang", "", "Language (go, java, ts...)")
	newCmd.Flags().StringVar(&newFramework, "framework", "", "Framework (gin, echo, spring, express...)")
	newCmd.Flags().StringVar(&newArch, "arch", "", "Code level architecture (layered, hexagonal, clean...)")
	newCmd.Flags().StringVar(&newProjectType, "type", "", "System level architecture (monolith, microservices...)")
	newCmd.Flags().StringVar(&newName, "name", "", "Name of the project to be created")

	rootCmd.AddCommand(newCmd)
}

func validateProjectInput(cat catalog.Catalog, input prompt.NewProjectInput) error {
	if input.Name == "" {
		return fmt.Errorf("project name required")
	}

	stack := catalog.Stack{
		Language:     input.Language,
		Framework:    input.Framework,
		ProjectType:  input.ProjectType,
		Architecture: input.Architecture,
	}

	if !cat.IsSupported(stack) {
		return fmt.Errorf("unsupported stack")
	}

	return nil
}
