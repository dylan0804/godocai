package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dylan0804/godocai/search"
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

type ViewState int

const (
	StateInput ViewState = iota
	StateResults
)

type AppModel struct {
	input   *InputModel
	results ResultsModel
	state   ViewState
	status  string
	width   int
	height  int
	loading bool
}

func NewAppModel() AppModel {
	return AppModel{
		input:   NewInputModel(),
		results: NewResultsModel(),
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
		m.status = ""
		m.input.textInput.Blur()
		m.state = StateResults
		if msg.err != nil {
			m.status = fmt.Sprintf("Error: %s", msg.err)
		} else {
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
	var content string

	header := TitleStyle.Render("Go Package Search") + "\n\n" +
			InputPromptStyle.Render() + m.input.View() + "\n\n"
	
	switch m.state {
	case StateInput:
		content = header + StatusMessageStyle(m.status)
	case StateResults:
		content = header + m.results.View()
	}

	return DocStyle.Render(content)
} 