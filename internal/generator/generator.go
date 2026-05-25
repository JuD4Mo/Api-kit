package generator

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
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

	basePath := filepath.Join(
		string(config.Language),
		string(config.Framework),
		string(config.ProjectType),
		string(config.Architecture),
	)

	// We use WalkDir in order to create folders and files
	err = fs.WalkDir(templates.FS, basePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// we need the relative path to the created project
		relativePath, err := filepath.Rel(basePath, path)
		if err != nil {
			return err
		}

		// root directory will be returned as .
		if relativePath == "." {
			return nil
		}

		outPath := filepath.Join(config.Name, relativePath)

		if d.IsDir() {
			return os.Mkdir(outPath, 0755)
		}

		if before, ok := strings.CutSuffix(outPath, ".tmpl"); ok {
			content, err := renderTemplate(path, data)
			if err != nil {
				return err
			}
			log.Println(before)
			return os.WriteFile(before, []byte(content), 0644)
		}

		return nil
	})

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
