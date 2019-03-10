package navigator

import (
	"fmt"
	"io"

	"github.com/gomarkdown/markdown/parser"
)

type ID string

type Page struct {
	ID    ID
	Title string
	File  string
	Subs  []ID
}

func (page Page) String() string {
	return fmt.Sprintf("%s (%q)", page.Title, page.File)
}

func ParsePage(re io.Reader) (Page, error) {

	return Page{}, nil
}

func DefaultParser() *parser.Parser {
	var parser = parser.New()
	return parser
}
