package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dylan0804/godocai/pkg/search"
)

type searchResultsMsg struct {
	results []list.Item
	err     error
}

func searchPackages(query string) tea.Cmd {
	return func() tea.Msg {
		results, err := search.Search(query)
		return searchResultsMsg{results, err}
	}
}

type AppModel struct {
	input   InputModel
	results ResultsModel
	status  string
	width   int
	height  int
	loading bool
}

func NewAppModel() AppModel {
	return AppModel{
		input:   NewInputModel(),
		results: NewResultsModel(),
		status:  "Type to search for Go packages",
	}
}

func (m AppModel) Init() tea.Cmd {
	return tea.Batch(
		m.input.Init(),
		m.results.Init(),
	)
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.loading || m.input.Value() == "" {
				return m, nil
			}
			
			m.loading = true
			m.status = "Searching..."
			
			return m, searchPackages(m.input.Value())
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		
		var cmd tea.Cmd
		m.input = m.input.SetWidth(m.width)
		m.results, cmd = m.results.Update(msg)
		cmds = append(cmds, cmd)

	case searchResultsMsg:
		m.loading = false
		if msg.err != nil {
			m.status = fmt.Sprintf("Error: %s", msg.err)
		} else {
			if len(msg.results) == 0 {
				m.status = "No results found"
			} else {
				m.status = fmt.Sprintf("Found %d results", len(msg.results))
			}
			m.results.results.SetItems(msg.results)
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	cmds = append(cmds, cmd)

	if _, ok := msg.(tea.WindowSizeMsg); !ok {
		m.results, cmd = m.results.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m AppModel) View() string {
	inputView := m.input.View()
	resultsView := m.results.View()
	statusView := StatusMessageStyle(m.status)

	return DocStyle.Render(
		TitleStyle.Render("Go Package Search") + "\n\n" +
		inputView + "\n" +
		resultsView + "\n" +
		statusView,
	)
} 