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

func getPackageInfo(link string) tea.Cmd {
	return func() tea.Msg {
		search.GetPackageInfo(link)
		fmt.Println("Package info fetched")
		return nil
	}
}

type ViewState int

const (
	StateInput ViewState = iota
	StateResults
	StateDetail
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
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		
		m.input = m.input.SetWidth(m.width)
		
		var cmd tea.Cmd
		m.results, cmd = m.results.Update(msg)
		return m, cmd
		
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
		
	case ItemSelectedMsg:
		m.state = StateDetail
		return m, nil
		
	case InputSubmitMsg:
		if m.loading || msg.Value == "" {
			return m, nil
		}
		
		m.loading = true
		m.status = "Searching..."
		return m, searchPackages(msg.Value)
		
	case searchResultsMsg:
		m.loading = false
		m.status = ""
		m.state = StateResults
		
		if msg.err != nil {
			m.status = fmt.Sprintf("Error: %s", msg.err)
		} else {
			m.results.results.SetItems(msg.results)
		}
		return m, getPackageInfo(msg.results[0].(search.Result).Link)
	}
	
	var cmd tea.Cmd
	
	switch m.state {
	case StateInput:
		m.input, cmd = m.input.Update(msg)
		return m, cmd
		
	case StateResults:
		// Only update results component when in results state
		m.results, cmd = m.results.Update(msg)
		return m, cmd
		
	case StateDetail:
		return m, getPackageInfo(m.results.results.SelectedItem().(search.Result).Link)
	}
	
	return m, nil
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
		content = fmt.Sprintln("Detail: aa")
	}

	return DocStyle.Render(content)
} 