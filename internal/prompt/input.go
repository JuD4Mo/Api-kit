package prompt

import "github.com/Jud4Mo/multi-templates/internal/catalog"

type NewProjectInput struct {
	Language     catalog.Language
	Framework    catalog.Framework
	ProjectType  catalog.ProjectType
	Architecture catalog.Architecture
	Name         string
}
