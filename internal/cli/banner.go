package cli

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	primary   = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4")).Bold(true)
	secondary = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575"))
	muted     = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262"))
)

func PrintBanner() {
	fmt.Println(primary.Render(`
             _       _    _ _   
  __ _ _ __ (_)     | | _(_) |_ 
 / _' | '_ \| |_____| |/ / | __|
| (_| | |_) | |_____|   <| | |_ 
 \__,_| .__/|_|     |_|\_\_|\__|
      |_|                       
`))
	fmt.Println(secondary.Render("  api-kit — project scaffolder"))
	fmt.Println(muted.Render("  version " + Version + " — answer a few questions and get a project ready"))
	fmt.Println()
}
