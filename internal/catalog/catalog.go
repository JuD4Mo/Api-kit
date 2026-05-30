package catalog

import "slices"

type (
	Language     string
	Framework    string
	ProjectType  string
	Architecture string
)

var Labels = map[string]string{
	"go":            " Go",
	"ts":            " TypeScript",
	"java":          " Java",
	"springboot":    " SpringBoot",
	"gin":           " Gin",
	"echo":          " Echo",
	"monolith":      " Monolith",
	"microservices": " Microservices",
	"layered":       " Layered",
	"hexagonal":     "󱃖 Hexagonal",
	"clean":         " Clean",
}

const (
	LanguageGo               Language     = "go"
	FrameworkGin             Framework    = "gin"
	ProjectTypeMonolith      ProjectType  = "monolith"
	ArchitectureLayered      Architecture = "layered"
	LanguageJava             Language     = "java"
	FrameworkSpringBoot      Framework    = "springboot"
	ProjectTypeMicroservices ProjectType  = "microservices"
)

type Stack struct {
	Language     Language
	Framework    Framework
	ProjectType  ProjectType
	Architecture Architecture
}

type Catalog struct {
	stacks []Stack
}

func NewCatalog() Catalog {
	return Catalog{
		stacks: []Stack{
			{
				Language:     LanguageGo,
				Framework:    FrameworkGin,
				ProjectType:  ProjectTypeMonolith,
				Architecture: ArchitectureLayered,
			},
			{
				Language:     LanguageJava,
				Framework:    FrameworkSpringBoot,
				ProjectType:  ProjectTypeMonolith,
				Architecture: ArchitectureLayered,
			},
		},
	}
}

func (c Catalog) Languages() []Language {
	languages := make([]Language, 0)

	for _, stack := range c.stacks {
		if slices.Contains(languages, stack.Language) {
			continue
		}

		languages = append(languages, stack.Language)
	}

	return languages
}

func (c Catalog) FrameworksByLanguage(lang Language) []Framework {
	frameworks := make([]Framework, 0)

	for _, stack := range c.stacks {
		if stack.Language != lang {
			continue
		}

		if slices.Contains(frameworks, stack.Framework) {
			continue
		}

		frameworks = append(frameworks, stack.Framework)
	}

	return frameworks
}

func (c Catalog) IsSupported(s Stack) bool {
	return slices.Contains(c.stacks, s)
}

func (c Catalog) ProjectTypesByLanguageAndFramework(lang Language, framework Framework) []ProjectType {
	projectTypes := make([]ProjectType, 0)

	for _, stack := range c.stacks {
		if stack.Language != lang || stack.Framework != framework {
			continue
		}

		if slices.Contains(projectTypes, stack.ProjectType) {
			continue
		}

		projectTypes = append(projectTypes, stack.ProjectType)
	}

	return projectTypes
}

func (c Catalog) ArchitecturesByBase(lang Language, framework Framework, pt ProjectType) []Architecture {
	archs := make([]Architecture, 0)

	for _, stack := range c.stacks {
		if stack.Language != lang || stack.Framework != framework || stack.ProjectType != pt {
			continue
		}

		if slices.Contains(archs, stack.Architecture) {
			continue
		}

		archs = append(archs, stack.Architecture)
	}

	return archs
}
