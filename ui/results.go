package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ResultsModel struct {
	results  list.Model
}

func NewDefaultDelegate() list.DefaultDelegate {
    d := list.DefaultDelegate{}
    d.ShowDescription = true
    
    // More contrasting color palette
    titleColor := lipgloss.Color("#36CFC9") 
    descColor := lipgloss.Color("#DDDDDD")     
    
	// normal items
    d.Styles.NormalTitle = lipgloss.NewStyle().
        Bold(true).
        Foreground(titleColor).
        MarginLeft(2)
    
    d.Styles.NormalDesc = lipgloss.NewStyle().
        Foreground(descColor).
        MarginLeft(4)
    
	// selected item
    d.Styles.SelectedTitle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#36CFC9")).
		Padding(0, 1).
		MarginLeft(1)

	d.Styles.SelectedDesc = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#000000")). 
		Background(lipgloss.Color("#36CFC9")).
		Padding(0, 1).
		MarginLeft(3)
    
    // dimensions
    d.SetHeight(2)
    d.SetSpacing(1)
    
    return d
}

func NewResultsModel() ResultsModel {
	d := NewDefaultDelegate()

	resultsList := list.New([]list.Item{}, d, 0, 0)
	resultsList.SetShowTitle(false)
	resultsList.SetShowStatusBar(false)
	resultsList.SetFilteringEnabled(false)

	return ResultsModel{
		results:  resultsList,
	}
}

func (m ResultsModel) Init() tea.Cmd {
	return nil
}

func (m ResultsModel) Update(msg tea.Msg) (ResultsModel, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.results.SetSize(msg.Width, msg.Height - 8)
	case tea.KeyMsg:
		if msg.Type == tea.KeyEnter {
			fmt.Println("Enter key pressed")
		}
	}
	
	m.results, cmd = m.results.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m ResultsModel) View() string {
	return m.results.View()
}
