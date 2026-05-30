package prompt

import "github.com/Jud4Mo/api-kit/internal/catalog"

type NewProjectInput struct {
	Language     catalog.Language
	Framework    catalog.Framework
	ProjectType  catalog.ProjectType
	Architecture catalog.Architecture
	Name         string
}

func NewProjectInputFromFlags(lang string, framework string, projectType string, arch string, name string) NewProjectInput {
	return NewProjectInput{
		Language:     catalog.Language(lang),
		Framework:    catalog.Framework(framework),
		ProjectType:  catalog.ProjectType(projectType),
		Architecture: catalog.Architecture(arch),
		Name:         name,
	}
}
