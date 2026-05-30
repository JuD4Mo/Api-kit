package prompt

import (
	"fmt"
	"strings"

	"github.com/Jud4Mo/multi-templates/internal/catalog"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7D56F4")).
			Bold(true).
			MarginBottom(1)

	langStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00ADD8")).
			Bold(true)

	frameworkStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575"))

	projectTypeStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#F39C12"))

	architectureStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#E84393"))

	treeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262"))
)

func PrintCatalog(cat catalog.Catalog) {
	fmt.Println(titleStyle.Render("󱚊 API-KIT Catalog"))

	for _, lang := range cat.Languages() {
		fmt.Println(langStyle.Render(FormatLabel(string(lang))))

		frameworks := cat.FrameworksByLanguage(lang)
		for i, fw := range frameworks {
			isLastFramework := i == len(frameworks)-1
			fwPrefix := " ├──"
			if isLastFramework {
				fwPrefix = " └──"
			}

			fmt.Printf("%s %s\n", treeStyle.Render(fwPrefix), frameworkStyle.Render(FormatLabel(string(fw))))

			projectTypes := cat.ProjectTypesByLanguageAndFramework(lang, fw)
			for j, pt := range projectTypes {
				isLastPt := j == len(projectTypes)-1

				indent := " │  "
				if isLastFramework {
					indent = "    "
				}

				ptPrefix := indent + " ├──"
				if isLastPt {
					ptPrefix = indent + " └──"
				}

				fmt.Printf("%s %s\n", treeStyle.Render(ptPrefix), projectTypeStyle.Render(FormatLabel(string(pt))))

				archs := cat.ArchitecturesByBase(lang, fw, pt)
				var archLabels []string
				for _, arch := range archs {
					archLabels = append(archLabels, FormatLabel(string(arch)))
				}
				archString := strings.Join(archLabels, ", ")

				archIndent := indent + " │  "
				if isLastPt {
					archIndent = indent + "    "
				}

				archPrefix := archIndent + "  └─ Code architectures: "
				fmt.Printf("%s%s\n", treeStyle.Render(archPrefix), architectureStyle.Render(archString))
			}
		}
		fmt.Println()
	}
}
