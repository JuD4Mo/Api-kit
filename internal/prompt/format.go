package prompt

import (
	"strings"

	"github.com/Jud4Mo/multi-templates/internal/catalog"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// FormatLabel takes a raw value (e.g. "go", "gin") and returns its pretty label with Nerd Fonts
// defined in catalog.Labels. If it's not found, it applies a title case fallback.
func FormatLabel(val string) string {
	if formatted, exists := catalog.Labels[strings.ToLower(val)]; exists {
		return formatted
	}

	caser := cases.Title(language.English)
	return caser.String(val)
}
