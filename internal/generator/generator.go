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
		"internal/controller",
		"internal/repository",
		"internal/service",
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
	return fmt.Sprintf("# %s\n", projectName)
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
