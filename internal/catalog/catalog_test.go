package catalog

import (
	"slices"
	"testing"
)

func TestLanguages(t *testing.T) {
	c := NewCatalog()

	got := c.Languages()
	want := []Language{LanguageGo, LanguageJava, LanguageRuby}

	if !slices.Equal(got, want) {
		t.Errorf("Languages don't match -> want: %v - got: %v", want, got)
	}
}

func TestFrameworksByLanguage(t *testing.T) {
	c := NewCatalog()
	lang := LanguageGo
	got := c.FrameworksByLanguage(lang)
	want := []Framework{FrameworkGin, FrameworkChi, FramewrokHttpNet}

	if !slices.Equal(got, want) {
		t.Errorf("Frameworks don't match -> want: %v - got: %v", want, got)
	}
}

func TestIsSupported(t *testing.T) {
	c := NewCatalog()

	stack := Stack{
		Language:     LanguageGo,
		Framework:    FrameworkGin,
		ProjectType:  ProjectTypeMonolith,
		Architecture: ArchitectureLayered,
	}

	if !c.IsSupported(stack) {
		t.Error("The stack is not supported")
	}
}

func TestProjectTypesByLanguageAndFramework(t *testing.T) {
	c := NewCatalog()

	got := c.ProjectTypesByLanguageAndFramework(LanguageGo, FrameworkGin)
	want := []ProjectType{ProjectTypeMonolith}

	if !slices.Equal(got, want) {
		t.Errorf("Project types don't match -> want: %v - got: %v", want, got)
	}
}

func TestArchitecturesByBase(t *testing.T) {
	c := NewCatalog()

	got := c.ArchitecturesByBase(LanguageGo, FrameworkGin, ProjectTypeMonolith)
	want := []Architecture{ArchitectureLayered}

	if !slices.Equal(got, want) {
		t.Errorf("Architecture types don't match -> want: %v - got: %v", want, got)
	}
}

func TestArchitecturesByBaseChi(t *testing.T) {
	c := NewCatalog()

	got := c.ArchitecturesByBase(LanguageGo, FrameworkChi, ProjectTypeMonolith)
	want := []Architecture{ArchitectureLayered, ArchitectureHexagonal, ArchitectureSriniously}

	if !slices.Equal(got, want) {
		t.Errorf("Architecture types don't match -> want: %v - got: %v", want, got)
	}
}

func TestIsSupportedNegative(t *testing.T) {
	c := NewCatalog()
	lang := Language("java")
	stack := Stack{
		Language:     lang,
		Framework:    FrameworkGin,
		ProjectType:  ProjectTypeMonolith,
		Architecture: ArchitectureLayered,
	}

	if c.IsSupported(stack) {
		t.Error("The stack should not be supported")
	}
}
