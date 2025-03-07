package search

import (
	"net/url"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/gocolly/colly"
)

type MethodInfo struct {
	Name        string
	Signature   string
	Description []string
}

type TypeInfo struct {
	Name        string
	Description string
	Methods     []MethodInfo
	Examples    []string
}

type Result struct {
    title       string
    Example     string
    Synopsis    string
	Link        string
}

func (r Result) Title() string { return r.title }
func (r Result) Description() string { return r.Example }
func (r Result) FilterValue() string { return r.Example }

func Search(term string) ([]list.Item, error) {
	results := []list.Item{}
	c := colly.NewCollector()

	c.OnHTML(".SearchSnippet", func(e *colly.HTMLElement) {
		rawTitle := e.ChildText(".SearchSnippet-headerContainer h2")
		link := e.ChildAttr(".SearchSnippet-headerContainer h2 a:first-of-type", "href")

		cleanTitle := strings.Join(strings.Fields(rawTitle), " ")

		synopsis := e.ChildText("p.SearchSnippet-infoLabel")

		example := e.ChildText("pre.SearchSnippet-symbolCode")

		results = append(results, Result{
			title:       cleanTitle,
			Synopsis:    synopsis,
			Example:     example,
			Link:        link,
		})
	})

	searchURL := "https://pkg.go.dev/search?q=" + url.QueryEscape(term) + "&m=symbol"
	err := c.Visit(searchURL)
	
	return results, err
} 

func GetPackageInfo(packageLink string, anchor string) (*TypeInfo, error) {
	c := colly.NewCollector()

	result := TypeInfo{}

	selector := "div.Documentation-type h4#" + anchor
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		result.Name = strings.TrimSpace(e.Text)
		result.Description = strings.TrimSpace(e.DOM.Next().Text())
	})

	c.Visit("https://pkg.go.dev/" + packageLink)
	
	return &TypeInfo{
		Name: result.Name,
	}, nil
}