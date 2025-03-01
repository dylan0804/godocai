package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type PackageDelegate struct {
	styles list.DefaultItemStyles
}

func (d PackageDelegate) Height() int { 
    return 10
}

func (d PackageDelegate) Spacing() int { 
    return 1 
}

func (d PackageDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { 
    return nil 
}

// func (d PackageDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
// 	item, ok := listItem.(search.Result)
// 	if !ok {
// 		return
// 	}

// 	styles := d.styles
// 	if index == m.Index() {
// 		styles = d.styles.selec
// 	}
// }
