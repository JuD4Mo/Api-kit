package prompt

import (
	"fmt"

	"github.com/Jud4Mo/api-kit/internal/catalog"
	"github.com/charmbracelet/huh"
)

// moduleLang is used to conditionally show the module input only for Go projects.
const moduleLang catalog.Language = "go"

/*
Generics use case:
using ~string means that we're allowing types whose underlying type is string
*/
func optionsFromValues[T ~string](values []T) []huh.Option[T] {
	options := make([]huh.Option[T], 0)
	for _, x := range values {
		label := FormatLabel(string(x))
		option := huh.NewOption(label, x)
		options = append(options, option)
	}

	return options
}

func AskNewProjectInput(cat catalog.Catalog, input NewProjectInput) (NewProjectInput, error) {
	if input.Language == "" {
		langOptions := optionsFromValues(cat.Languages())
		selectLanguageHuh := huh.NewSelect[catalog.Language]().
			Title("Which language do you want to use?").
			Options(langOptions...).
			Value(&input.Language)

		err := huh.NewForm(
			huh.NewGroup(
				selectLanguageHuh,
			),
		).Run()
		if err != nil {
			return input, fmt.Errorf("run new project prompt: %w", err)
		}
	}

	if input.Framework == "" {
		frameworkOptions := optionsFromValues(
			cat.FrameworksByLanguage(input.Language),
		)

		selectFrameworkHuh := huh.NewSelect[catalog.Framework]().
			Title("Which framework do you want to use?").
			Options(frameworkOptions...).
			Value(&input.Framework)

		err := huh.NewForm(
			huh.NewGroup(
				selectFrameworkHuh,
			),
		).Run()
		if err != nil {
			return input, fmt.Errorf("run new project prompt: %w", err)
		}
	}

	if input.ProjectType == "" {
		typesOptions := optionsFromValues(
			cat.ProjectTypesByLanguageAndFramework(input.Language, input.Framework),
		)
		selectProjectTypeHuh := huh.NewSelect[catalog.ProjectType]().
			Title("Which system architecture do you want to implement?").
			Options(typesOptions...).
			Value(&input.ProjectType)

		err := huh.NewForm(
			huh.NewGroup(
				selectProjectTypeHuh,
			),
		).Run()
		if err != nil {
			return input, fmt.Errorf("run new project prompt: %w", err)
		}
	}

	if input.Architecture == "" {
		archsOptions := optionsFromValues(
			cat.ArchitecturesByBase(input.Language, input.Framework, input.ProjectType),
		)
		selectArchOptions := huh.NewSelect[catalog.Architecture]().
			Title("Which code architecture do you want to use?").
			Options(archsOptions...).
			Value(&input.Architecture)

		err := huh.NewForm(
			huh.NewGroup(
				selectArchOptions,
			),
		).Run()
		if err != nil {
			return input, fmt.Errorf("run new project prompt: %w", err)
		}
	}

	if input.Name == "" {
		projectName := huh.NewInput().
			Title("What is the project name?").
			Value(&input.Name).
			Validate(validateName)

		err := huh.NewForm(
			huh.NewGroup(
				projectName,
			),
		).Run()
		if err != nil {
			return input, fmt.Errorf("run new project prompt: %w", err)
		}
	}

	if input.Module == "" && input.Language == moduleLang {
		moduleInput := huh.NewInput().
			Title("What is the Go module path?").
			Description("e.g. github.com/username/project").
			Value(&input.Module)

		err := huh.NewForm(
			huh.NewGroup(
				moduleInput,
			),
		).Run()
		if err != nil {
			return input, fmt.Errorf("run new project prompt: %w", err)
		}
	}

	return input, nil
}

func validateName(projectName string) error {
	if projectName == "" {
		return fmt.Errorf("name can not be empty")
	}
	return nil
}
