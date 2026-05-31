package generator

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/Jud4Mo/api-kit/internal/catalog"
)

func TestGenerate(t *testing.T) {
	// temporal directory
	tmp := t.TempDir()
	completePath := filepath.Join(tmp, "something")

	config := ProjectConfig{
		Language:     catalog.LanguageGo,
		Framework:    catalog.FrameworkGin,
		ProjectType:  catalog.ProjectTypeMonolith,
		Architecture: catalog.ArchitectureLayered,
		Name:         completePath,
		Module:       "something",
	}

	err := Generate(config)
	if err != nil {
		t.Fatalf("failed: %v", err)
	}

	readme := filepath.Join(completePath, "README.md")
	gomod := filepath.Join(completePath, "go.mod")
	maingo := filepath.Join(completePath, "cmd", "api", "main.go")
	entity := filepath.Join(completePath, "internal", "user", "entity.go")
	controller := filepath.Join(completePath, "internal", "user", "endpoint.go")
	repo := filepath.Join(completePath, "internal", "user", "repository.go")
	service := filepath.Join(completePath, "internal", "user", "service.go")
	boots := filepath.Join(completePath, "pkg", "bootstrap", "bootstrap.go")

	assertPathExists(t, readme)
	assertPathExists(t, gomod)
	assertPathExists(t, maingo)
	assertPathExists(t, entity)
	assertPathExists(t, controller)
	assertPathExists(t, repo)
	assertPathExists(t, service)
	assertPathExists(t, boots)

}

func TestGenerateError(t *testing.T) {
	config := ProjectConfig{
		Language:     catalog.LanguageGo,
		Framework:    catalog.FrameworkGin,
		ProjectType:  catalog.ProjectTypeMonolith,
		Architecture: catalog.ArchitectureLayered,
		Name:         "",
	}

	err := Generate(config)
	if err == nil {
		t.Fatal("Generate() error = nil, want error")
	}

}

func TestGenerateE2E(t *testing.T) {
	tmp := t.TempDir()
	completePath := filepath.Join(tmp, "something")

	config := ProjectConfig{
		Language:     catalog.LanguageGo,
		Framework:    catalog.FrameworkGin,
		ProjectType:  catalog.ProjectTypeMonolith,
		Architecture: catalog.ArchitectureLayered,
		Name:         completePath,
		Module:       "something",
	}

	err := Generate(config)
	if err != nil {
		t.Fatalf("failed: %v", err)
	}

	readme := filepath.Join(completePath, "README.md")
	gomod := filepath.Join(completePath, "go.mod")
	maingo := filepath.Join(completePath, "cmd", "api", "main.go")
	entity := filepath.Join(completePath, "internal", "user", "entity.go")
	controller := filepath.Join(completePath, "internal", "user", "endpoint.go")
	repo := filepath.Join(completePath, "internal", "user", "repository.go")
	service := filepath.Join(completePath, "internal", "user", "service.go")
	boots := filepath.Join(completePath, "pkg", "bootstrap", "bootstrap.go")

	assertPathExists(t, readme)
	assertPathExists(t, gomod)
	assertPathExists(t, maingo)
	assertPathExists(t, entity)
	assertPathExists(t, controller)
	assertPathExists(t, repo)
	assertPathExists(t, service)
	assertPathExists(t, boots)

	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)

	// ensures execution
	projectDir := completePath
	err = os.Chdir(projectDir)
	if err != nil {
		t.Errorf("failed to change to generated project directory: %v", err)
	}

	cmdTidy := exec.Command("go", "mod", "tidy")
	cmdTidy.Stdout = os.Stdout
	cmdTidy.Stderr = os.Stderr

	err = cmdTidy.Run()
	if err != nil {
		t.Errorf("failed to run 'go mod tidy': %v", err)
	}

	cmdBuild := exec.Command("go", "build", "./cmd/api")
	cmdBuild.Stdout = os.Stdout
	cmdBuild.Stderr = os.Stderr
	err = cmdBuild.Run()
	if err != nil {
		t.Fatalf("Failed to run 'go build': %v", err)
	}
}

func assertPathExists(t *testing.T, path string) {
	t.Helper()
	_, err := os.Stat(path)
	if err != nil {
		t.Fatalf("failed finding path: %v", err)
	}
}
