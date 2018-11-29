package lib

import (
	"fmt"

	"github.com/upamune/go-esa/esa"
)

// DisplayConsole is console struct
type DisplayConsole struct {
}

// ShowArticles is print article
func (d *DisplayConsole) ShowArticles(articles []esa.PostResponse) error {
	for _, value := range articles {
		fmt.Println(value.Name)
	}

	return nil
}
