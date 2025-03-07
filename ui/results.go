package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dylan0804/godocai/search"
)

type ResultsModel struct {
	results  list.Model
}

type ItemSelectedMsg struct {
	Item search.Result
}

func NewDefaultDelegate() list.DefaultDelegate {
    d := list.DefaultDelegate{}
    d.ShowDescription = true
    
	// normal items
	d.Styles.NormalTitle = lipgloss.NewStyle().
        Bold(true).
        Foreground(PrimaryColor).
        BorderLeft(true).
        BorderStyle(lipgloss.NormalBorder()).
        BorderForeground(SubtleColor).
        PaddingLeft(1)
    
    d.Styles.NormalDesc = lipgloss.NewStyle().
        Foreground(SecondaryColor).
        BorderLeft(true).
        BorderStyle(lipgloss.NormalBorder()).
        BorderForeground(SubtleColor).
        PaddingLeft(1)
    
	// selected items
	d.Styles.SelectedTitle = lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("#FFFFFF")).
        Background(SelectedBg).
        BorderLeft(true).
        BorderStyle(lipgloss.NormalBorder()).
        BorderForeground(PrimaryColor).
        Padding(0, 1, 0, 1)
    
    d.Styles.SelectedDesc = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#EEEEEE")).
        Background(SelectedBg).
        BorderLeft(true).
        BorderStyle(lipgloss.NormalBorder()).
        BorderForeground(PrimaryColor).
        Padding(0, 1, 0, 1)
    
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
		switch msg.String() {
		case "enter":
			if item, ok := m.results.SelectedItem().(search.Result); ok {
				return m, HandleItemSelected(item)
			}
		}
	case SearchResultsMsg:
		m.results.SetItems(msg.results)
		return m, m.ChangeState(StateResults)
	}
	
	m.results, cmd = m.results.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}


func (m ResultsModel) View() string {
	return m.results.View()
}

func HandleItemSelected(item search.Result) tea.Cmd {
	return func() tea.Msg {
		return ItemSelectedMsg{Item: item}
	}
}

func (m ResultsModel) ChangeState(state ViewState) tea.Cmd {
	return func() tea.Msg {
		return StateChangeMsg{State: state}
	}
}

