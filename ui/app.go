package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dylan0804/godocai/ai"
	"github.com/dylan0804/godocai/search"
)
type packageInfoMsg struct {
	info search.TypeInfo
	err  error
}

func getPackageInfo(link string, anchor string) tea.Cmd {
	return func() tea.Msg {
		result, err := search.GetPackageInfo(link, anchor)
		return packageInfoMsg{*result, err}
	}
}

type ViewState int

const (
	StateInput ViewState = iota
	StateResults
	StateDetail
	StateLoading
)

type StateChangeMsg struct {
	State ViewState
}

type AppModel struct {
	input   *InputModel
	results ResultsModel
	detail  *DetailModel
	state   ViewState
	status  string
	width   int
	height  int
}

func NewAppModel() AppModel {
	return AppModel{
		input:   NewInputModel(),
		results: NewResultsModel(),
		detail:  NewDetailModel(),
		state:   StateInput,
	}
}

func (m AppModel) Init() tea.Cmd {
	return tea.Batch(
		m.input.Init(),
		m.results.Init(),
	)
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		m.detail.SetWidth(m.width)
		m.detail.SetHeight(m.height - 8)

		m.detail.SetViewportWidth(m.width)
		m.detail.SetViewportHeight(m.height - 8)

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	case ItemSelectedMsg:
		link := strings.SplitAfter(msg.Item.Link, "#")[1]
		return m, getPackageInfo(msg.Item.Link, link)

	case packageInfoMsg:
		m.detail.ready = true
		m.state = StateDetail

		if msg.err != nil {
			m.status = fmt.Sprintf("Error: %s", msg.err)
		} else {
			m.detail.info = msg.info
			fmt.Println(msg.info.Description)
			return m, ai.StreamAIExplanation(msg.info.Description)
		}
	
	case StateChangeMsg:
		m.state = msg.State
		return m, nil
	}
		
	m.input, cmd = m.input.Update(msg)
	cmds = append(cmds, cmd)

	m.results, cmd = m.results.Update(msg)
	cmds = append(cmds, cmd)

	m.detail, cmd = m.detail.Update(msg)
	cmds = append(cmds, cmd)
	
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
	case StateDetail:
		content = m.detail.View()
	case StateLoading:
		content = header + StatusMessageStyle("Loading...")
	}

	return DocStyle.Render(content)
} 