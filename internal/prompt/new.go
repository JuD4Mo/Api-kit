package prompt

import (
	"fmt"

	"github.com/Jud4Mo/multi-templates/internal/catalog"
	"github.com/charmbracelet/huh"
)

/*
Generics use case:
using ~string means that we're allowing types whose underlying type is string
*/
func optionsFromValues[T ~string](values []T) []huh.Option[T] {
	options := make([]huh.Option[T], 0)
	for _, x := range values {
		label := string(x)
		option := huh.NewOption(label, x)
		options = append(options, option)
	}

	return options
}

func AskNewProjectInput(cat catalog.Catalog) (NewProjectInput, error) {
	newProjectInput := NewProjectInput{}
	langOptions := optionsFromValues(cat.Languages())
	selectHuh := huh.NewSelect[catalog.Language]().
		Title("Which language do you want to use?").
		Options(langOptions...).
		Value(&newProjectInput.Language)

	err := huh.NewForm(
		huh.NewGroup(
			selectHuh,
		),
	).Run()
	if err != nil {
		return newProjectInput, fmt.Errorf("run new project prompt: %w", err)
	}

	return newProjectInput, nil
}
