package search

import (
	"net/url"

	"github.com/charmbracelet/bubbles/list"
	"github.com/gocolly/colly"
)

type Result struct {
	title       string
	PackagePath string
	Synopsis    string
	URL         string
}

func (r Result) Title() string { return r.title }
func (r Result) Description() string { return r.Synopsis }
func (r Result) FilterValue() string { return r.PackagePath }

func Search(term string) ([]list.Item, error) {
	results := []list.Item{}
	c := colly.NewCollector()

	c.OnHTML(".SearchSnippet", func(e *colly.HTMLElement) {
		title := e.ChildText("h2")
		path := e.ChildText(".SearchSnippet-header-path")
		synopsis := e.ChildText(".SearchSnippet-synopsis")

		if title != "" {
			results = append(results, Result{
				title:       title,
				PackagePath: path,
				Synopsis:    synopsis,
			})
		}
	})

	searchURL := "https://pkg.go.dev/search?q=" + url.QueryEscape(term)
	err := c.Visit(searchURL)
	
	return results, err
} 