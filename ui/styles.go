package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	PrimaryColor = lipgloss.Color("#9D50BB")
    SecondaryColor = lipgloss.Color("#EEEEEE")
    SubtleColor = lipgloss.Color("#666666")
    SelectedBg = lipgloss.Color("#6A3093")
    

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

	// detail
	TypeNameStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#9D50BB")).
		BorderBottom(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#6A3093")).
		MarginBottom(1).
		Padding(0, 0, 0, 0)

	DescriptionStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#EEEEEE")).
		MarginBottom(1).
		Padding(0, 1)

	SectionTitleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#9D50BB")).
		PaddingBottom(1)

	MethodNameStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#BBBBBB"))

	SignatureStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#DDDDDD")).
		Background(lipgloss.Color("#333333")).
		Padding(1).
		MarginTop(1).
		MarginBottom(1)

	ExampleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#DDDDDD")).
		Background(lipgloss.Color("#333333")).
		Padding(1).
		MarginTop(1).
		MarginBottom(1)
) 