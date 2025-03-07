package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dylan0804/godocai/search"
)

type InputModel struct {
	textInput textinput.Model
	width     int
}

type SearchResultsMsg struct {
	results []list.Item
}

func NewInputModel() *InputModel {
	ti := textinput.New()
	ti.Placeholder = "Enter package to search..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 40

	return &InputModel{
		textInput: ti,
		width:     40,
	}
}

func (m *InputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *InputModel) Update(msg tea.Msg) (*InputModel, tea.Cmd) {
	var (
		cmd tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m, cmd = m.SearchPackage()
			cmds = append(cmds, cmd)
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *InputModel) View() string {
	return InputPromptStyle.Render("Search: ") + m.textInput.View()
}

func (m *InputModel) Value() string {
	return m.textInput.Value()
}

func (m *InputModel) SetWidth(width int) *InputModel {
	m.width = width
	m.textInput.Width = width - 10 
	return m
}

func (m *InputModel) SearchPackage() (*InputModel, tea.Cmd) {
	results, err := search.Search(m.Value())
	if err != nil {
		return m, nil
	}

	m.textInput.Blur()

	return m, m.HandleSearchResults(results)
}

func (m *InputModel) HandleSearchResults(results []list.Item) tea.Cmd {
	return func() tea.Msg {
		return SearchResultsMsg{results: results}
	}
}

func (m *InputModel) ChangeState(state ViewState) tea.Cmd {
	return func() tea.Msg {
		return StateChangeMsg{State: state}
	}
}
