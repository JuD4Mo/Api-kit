package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Jud4Mo/multi-templates/internal/catalog"
)

type ProjectConfig struct {
	Language     catalog.Language
	Framework    catalog.Framework
	ProjectType  catalog.ProjectType
	Architecture catalog.Architecture
	Name         string
}

func Generate(config ProjectConfig) error {
	if config.Name == "" {
		return fmt.Errorf("project name can not be empty")
	}

	err := os.Mkdir(config.Name, 0755)
	if err != nil {
		return fmt.Errorf("create project directory: %w", err)
	}

	readmePath := "README.md"
	content := readmeContent(config.Name)
	err = writeProjectFile(config.Name, readmePath, content)
	if err != nil {
		return err
	}

	modulePath := "go.mod"
	content = goModContent(config.Name)
	err = writeProjectFile(config.Name, modulePath, content)
	if err != nil {
		return err
	}

	paths := []string{
		"cmd/api",
		"internal/item",
		"internal/domain",
		"pkg/bootstrap",
	}

	for _, relativePath := range paths {
		completePath := filepath.Join(config.Name, relativePath)
		err := os.MkdirAll(completePath, 0755)
		if err != nil {
			return fmt.Errorf("generating structure error: %w", err)
		}
	}

	mainPath := filepath.Join("cmd", "api", "main.go")
	content = mainGoContent()
	err = writeProjectFile(config.Name, mainPath, content)
	if err != nil {
		return err
	}

	bootstrapPath := filepath.Join("pkg", "bootstrap", "bootstrap.go")
	content = bootstrapContent()
	err = writeProjectFile(config.Name, bootstrapPath, content)
	if err != nil {
		return err
	}

	itemPath := filepath.Join("internal", "domain", "item.go")
	content = itemContent()
	err = writeProjectFile(config.Name, itemPath, content)
	if err != nil {
		return err
	}

	repositoryPath := filepath.Join("internal", "item", "repository.go")
	content = repositoryContent()
	err = writeProjectFile(config.Name, repositoryPath, content)
	if err != nil {
		return err
	}

	servicePath := filepath.Join("internal", "item", "service.go")
	content = serviceContent()
	err = writeProjectFile(config.Name, servicePath, content)
	if err != nil {
		return err
	}

	controllerPath := filepath.Join("internal", "item", "controller.go")
	content = controllerContent()
	err = writeProjectFile(config.Name, controllerPath, content)
	if err != nil {
		return err
	}

	return nil
}

func writeProjectFile(projectName string, relativePath string, content string) error {
	completePath := filepath.Join(projectName, relativePath)
	err := os.WriteFile(completePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("writing %s: %w", completePath, err)
	}
	return nil
}

func readmeContent(projectName string) string {
	return fmt.Sprintf(`
# %s

## Executes this following commands
go mod tidy

## Start web server
go run ./cmd/api
`, projectName)
}

func goModContent(projectName string) string {
	return fmt.Sprintf("module %s\n\ngo 1.22\n", projectName)
}

func mainGoContent() string {
	return `package main

import "fmt"

func main() {
	fmt.Println("working")
}
`
}

func bootstrapContent() string {
	return `
	
`
}

func itemContent() string {
	return `
	
`
}

func repositoryContent() string {
	return `
	
`
}

func serviceContent() string {
	return `
	
`
}

func controllerContent() string {
	return `
	
`
}
