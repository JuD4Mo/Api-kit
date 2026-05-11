package generator

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Jud4Mo/multi-templates/internal/catalog"
	"github.com/Jud4Mo/multi-templates/internal/templates"
)

type ProjectConfig struct {
	Language     catalog.Language
	Framework    catalog.Framework
	ProjectType  catalog.ProjectType
	Architecture catalog.Architecture
	Name         string
	Module       string
}

type templateData struct {
	ProjectName string
	Module      string
}

func Generate(config ProjectConfig) error {
	if config.Name == "" {
		return fmt.Errorf("project name can not be empty")
	}
	if config.Module == "" {
		config.Module = config.Name
	}

	err := os.Mkdir(config.Name, 0755)
	if err != nil {
		return fmt.Errorf("create project directory: %w", err)
	}

	data := templateData{
		ProjectName: config.Name,
		Module:      config.Module,
	}

	//Render README.md file
	readmePath := "README.md"
	content, err := renderTemplate("go/gin/monolith/layered/README.md.tmpl", data)
	if err != nil {
		return err
	}

	err = writeProjectFile(config.Name, readmePath, content)
	if err != nil {
		return err
	}

	//Render go.mod file
	modulePath := "go.mod"
	content, err = renderTemplate("go/gin/monolith/layered/go.mod.tmpl", data)
	if err != nil {
		return err
	}

	err = writeProjectFile(config.Name, modulePath, content)
	if err != nil {
		return err
	}

	//Define project paths
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

	//Render main.go file
	mainPath := filepath.Join("cmd", "api", "main.go")
	content, err = renderTemplate("go/gin/monolith/layered/cmd/api/main.go.tmpl", data)
	if err != nil {
		return err
	}

	err = writeProjectFile(config.Name, mainPath, content)
	if err != nil {
		return err
	}

	//Render bootstrap.go file
	bootstrapPath := filepath.Join("pkg", "bootstrap", "bootstrap.go")
	content, err = renderTemplate("go/gin/monolith/layered/pkg/bootstrap/bootstrap.go.tmpl", data)
	if err != nil {
		return err
	}

	err = writeProjectFile(config.Name, bootstrapPath, content)
	if err != nil {
		return err
	}

	//Render item.go file (domain)
	itemPath := filepath.Join("internal", "domain", "item.go")
	content, err = renderTemplate("go/gin/monolith/layered/internal/domain/item.go.tmpl", data)
	if err != nil {
		return err
	}

	err = writeProjectFile(config.Name, itemPath, content)
	if err != nil {
		return err
	}

	//Render repository.go file
	repositoryPath := filepath.Join("internal", "item", "repository.go")
	content, err = renderTemplate("go/gin/monolith/layered/internal/item/repository.go.tmpl", data)
	if err != nil {
		return err
	}

	err = writeProjectFile(config.Name, repositoryPath, content)
	if err != nil {
		return err
	}

	//Render service.go file
	servicePath := filepath.Join("internal", "item", "service.go")
	content, err = renderTemplate("go/gin/monolith/layered/internal/item/service.go.tmpl", data)
	if err != nil {
		return err
	}

	err = writeProjectFile(config.Name, servicePath, content)
	if err != nil {
		return err
	}

	//Render controller.go file
	controllerPath := filepath.Join("internal", "item", "controller.go")
	content, err = renderTemplate("go/gin/monolith/layered/internal/item/controller.go.tmpl", data)
	if err != nil {
		return err
	}

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

func renderTemplate(path string, data templateData) (string, error) {
	content, err := templates.FS.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read template %s: %w", path, err)
	}

	tmpl, err := template.New(filepath.Base(path)).Parse(string(content))
	if err != nil {
		return "", fmt.Errorf("parse template %s: %w", path, err)
	}

	var rendered bytes.Buffer

	err = tmpl.Execute(&rendered, data)
	if err != nil {
		return "", fmt.Errorf("render template %s: %w", path, err)
	}

	return rendered.String(), nil
}
