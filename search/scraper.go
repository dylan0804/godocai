package search

import (
	"net/url"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/gocolly/colly"
)

type Result struct {
    title       string
    Example     string
    Synopsis    string
}

func (r Result) Title() string { return r.title }
func (r Result) Description() string { return r.Example }
func (r Result) FilterValue() string { return r.Example }

func Search(term string) ([]list.Item, error) {
	results := []list.Item{}
	c := colly.NewCollector()

	c.OnHTML(".SearchSnippet", func(e *colly.HTMLElement) {
		rawTitle := e.ChildText(".SearchSnippet-headerContainer h2")

		cleanTitle := strings.Join(strings.Fields(rawTitle), " ")

		synopsis := e.ChildText("p.SearchSnippet-infoLabel")

		example := e.ChildText("pre.SearchSnippet-symbolCode")

		results = append(results, Result{
			title:       cleanTitle,
			Synopsis:    synopsis,
			Example:     example,
		})
	})

	searchURL := "https://pkg.go.dev/search?q=" + url.QueryEscape(term) + "&m=symbol"
	err := c.Visit(searchURL)
	
	return results, err
} 