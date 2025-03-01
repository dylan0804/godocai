package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ResultsModel struct {
	results  list.Model
}

func NewResultsModel() ResultsModel {
	return ResultsModel{
		results:  list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
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
		fmt.Println(msg.Width, msg.Height)
		m.results.SetSize(msg.Width - 4, msg.Height - 8)
	}
	
	m.results, cmd = m.results.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m ResultsModel) View() string {
	return m.results.View()
}
