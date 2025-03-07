package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dylan0804/godocai/ai"
	"github.com/dylan0804/godocai/search"
)

var Program *tea.Program

type DetailModel struct {
	info search.TypeInfo
	width int
	height int
	viewport viewport.Model
	ready bool
}

func NewDetailModel() *DetailModel {
	return &DetailModel{
		info: search.TypeInfo{},
		ready: false,
	}
}

func (m *DetailModel) Init() tea.Cmd {
	return nil
}

func (m *DetailModel) Update(msg tea.Msg) (*DetailModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
		switch msg.String() {
		case "up", "k":
			m.viewport.LineUp(1)
		case "down", "j":
			m.viewport.LineDown(1)
		case "pgup":
			m.viewport.HalfViewUp()
		case "pgdown":
			m.viewport.HalfViewDown()
		case "home":
			m.viewport.GotoTop()
		case "end":
			m.viewport.GotoBottom()
		}

	case ai.AIStreamMsg:
		m.info.Description = msg.Content
		m.viewport.SetContent(m.RenderContent())
		return m, nil
	}

	m.viewport.SetContent(m.RenderContent())
	m.viewport, cmd = m.viewport.Update(msg)

	return m, cmd
}

func (m *DetailModel) RenderContent() string {
	var sb strings.Builder

	sb.WriteString(TypeNameStyle.Render(fmt.Sprintf("Type: %s", m.info.Name)))
	sb.WriteString("\n\n")
	sb.WriteString(DescriptionStyle.Render(m.info.Description))
	sb.WriteString("\n\n")

	return sb.String()
}

func (m *DetailModel) View() string {
	if !m.ready {
		return "Loading..."
	}

	return fmt.Sprintf("%s\n%s", 
		TypeNameStyle.Render("Go Package Documentation"),
		m.viewport.View(),
	)
}

func(m *DetailModel) SetWidth(width int) {
	m.width = width
}

func(m *DetailModel) SetHeight(height int) {
	m.height = height
}

func(m *DetailModel) SetViewportWidth(width int) {
	m.viewport.Width = width
}

func(m *DetailModel) SetViewportHeight(height int) {
	m.viewport.Height = height
}
