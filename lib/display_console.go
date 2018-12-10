package lib

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/upamune/go-esa/esa"
)

// DisplayConsole is console struct
type DisplayConsole struct {
}

// ShowTimeStamp is output colored timestamp
func (d *DisplayConsole) ShowTimeStamp(timestamp string) string {
	green := color.New(color.FgGreen).SprintFunc()
	return green(timestamp)
}

// ShowUrl is output colored url
func (d *DisplayConsole) ShowUrl(url string) string {
	yellow := color.New(color.FgYellow).SprintFunc()
	return yellow(url)
}

// ShowArticles is print article
func (d *DisplayConsole) ShowArticles(article esa.PostResponse) error {
	createdAt := article.CreatedAt

	fmt.Fprintf(color.Output, "%s\t%s\t%s\n", d.ShowUrl(article.URL), d.ShowTimeStamp(createdAt), article.FullName)
	return nil
}
