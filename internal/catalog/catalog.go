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
	"ts":            " TypeScript",
	"js":            " JavaScript",
	"java":          " Java",
	"nestjs":        " NestJS",
	"springboot":    " Spring Boot",
	"gin":           " Gin",
	"echo":          " Echo",
	"express":       " Express",
	"monolith":      " Monolith",
	"microservices": " Microservices",
	"layered":       " Layered",
	"hexagonal":     "󱃖 Hexagonal",
	"sriniously":    "󰛨 Hexagonal Sriniously-based",
	"clean":         " Clean",
	"ruby":          " Ruby",
	"rails":         "󰫏 Rails",
	"python":        " Python",
	"django":        " Django",
	"http/net":      " http/net",
}

const (
	LanguageGo               Language     = "go"
	FrameworkGin             Framework    = "gin"
	ProjectTypeMonolith      ProjectType  = "monolith"
	ArchitectureLayered      Architecture = "layered"
	LanguageJava             Language     = "java"
	FrameworkSpringBoot      Framework    = "springboot"
	ProjectTypeMicroservices ProjectType  = "microservices"
	LanguageTs               Language     = "ts"
	FrameworkNestJS          Framework    = "nestjs"
	FrameworkExpress         Framework    = "express"
	LanguageRuby             Language     = "ruby"
	FrameworkRails           Framework    = "rails"
	LanguagePython           Language     = "python"
	FrameworkDjango          Framework    = "django"
	FrameworkEcho            Framework    = "echo"
	FrameworkChi             Framework    = "chi"
	FramewrokHttpNet         Framework    = "http/net"
	ArchitectureHexagonal    Architecture = "hexagonal"
	ArchitectureSriniously   Architecture = "sriniously"
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
				Language:     LanguageGo,
				Framework:    FrameworkChi,
				ProjectType:  ProjectTypeMonolith,
				Architecture: ArchitectureLayered,
			},
			{
				Language:     LanguageGo,
				Framework:    FrameworkChi,
				ProjectType:  ProjectTypeMonolith,
				Architecture: ArchitectureHexagonal,
			},
			{
				Language:     LanguageGo,
				Framework:    FrameworkChi,
				ProjectType:  ProjectTypeMonolith,
				Architecture: ArchitectureSriniously,
			},
			{
				Language:     LanguageGo,
				Framework:    FramewrokHttpNet,
				ProjectType:  ProjectTypeMonolith,
				Architecture: ArchitectureLayered,
			},
			{
				Language:     LanguageGo,
				Framework:    FramewrokHttpNet,
				ProjectType:  ProjectTypeMonolith,
				Architecture: ArchitectureHexagonal,
			},
			{
				Language:     LanguageJava,
				Framework:    FrameworkSpringBoot,
				ProjectType:  ProjectTypeMonolith,
				Architecture: ArchitectureLayered,
			},
			{
				Language:     LanguageGo,
				Framework:    FrameworkEcho,
				ProjectType:  ProjectTypeMonolith,
				Architecture: ArchitectureLayered,
			},
			{
				Language:     LanguageGo,
				Framework:    FrameworkEcho,
				ProjectType:  ProjectTypeMonolith,
				Architecture: ArchitectureHexagonal,
			},
			{
				Language:     LanguageGo,
				Framework:    FrameworkGin,
				ProjectType:  ProjectTypeMonolith,
				Architecture: ArchitectureHexagonal,
			},
			{
				Language:     LanguageRuby,
				Framework:    FrameworkRails,
				ProjectType:  ProjectTypeMonolith,
				Architecture: ArchitectureLayered,
			},
			{
				Language:     LanguageTs,
				Framework:    FrameworkNestJS,
				ProjectType:  ProjectTypeMonolith,
				Architecture: ArchitectureLayered,
			},
			{
				Language:     LanguageTs,
				Framework:    FrameworkNestJS,
				ProjectType:  ProjectTypeMicroservices,
				Architecture: ArchitectureLayered,
			},
			{
				Language:     LanguageTs,
				Framework:    FrameworkExpress,
				ProjectType:  ProjectTypeMonolith,
				Architecture: ArchitectureLayered,
			},
			{
				Language:     LanguagePython,
				Framework:    FrameworkDjango,
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
