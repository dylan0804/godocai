package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type InputModel struct {
	textInput textinput.Model
	width     int
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
	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
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