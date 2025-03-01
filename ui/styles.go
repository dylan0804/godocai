package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	DocStyle = lipgloss.NewStyle().Margin(1, 2)

	TitleStyle = lipgloss.NewStyle().
			Foreground(highlight).
			Bold(true).
			MarginBottom(1)

	InputPromptStyle = lipgloss.NewStyle().
			Foreground(special).
			Bold(true)

	InputStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(subtle).
			PaddingLeft(2)

	ResultTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(highlight)

	ResultPathStyle = lipgloss.NewStyle().
			Foreground(special).
			Italic(true)

	ResultSynopsisStyle = lipgloss.NewStyle().
			MarginLeft(2)

	StatusMessageStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}).
			Render

	SelectedResultStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FFFFFF"})
) 