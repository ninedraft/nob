package navigator

import (
	"bytes"
	"fmt"
	"io"
	"path"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

type ID string

func ExtractIDFromFileName(fileName string) ID {
	var base = path.Base(fileName)
	var ext = path.Ext(fileName)
	return ID(base[:len(base)-len(ext)])
}

type Page struct {
	ID    ID
	Title string
	File  string
	Subs  []ID
}

func (page Page) String() string {
	return fmt.Sprintf("%s (%q)", page.Title, page.File)
}

// TODO: (twentydraft) fix parser: TestParser works, but TestParsePage doesn't pass!
func ParsePage(fileName string, re io.Reader) (Page, error) {
	var buf = &bytes.Buffer{}
	if _, err := buf.ReadFrom(buf); err != nil {
		return Page{}, err
	}
	var mdAST = DefaultParser().Parse(buf.Bytes())
	var subPages []ID
	var title string
	var walker = func(node ast.Node, entering bool) ast.WalkStatus {
		switch node := node.(type) {
		case *ast.Link:
			var link = string(node.Destination)
			var ext = path.Ext(link)
			if !path.IsAbs(link) && ext == ".md" {
				subPages = append(subPages, ExtractIDFromFileName(link))
			}
		case *ast.Heading:
			var text = node.GetChildren()[0].(*ast.Text)
			title = string(text.Literal)
		default:

		}
		return ast.GoToNext
	}
	ast.Walk(mdAST, ast.NodeVisitorFunc(walker))
	return Page{
		ID:    ExtractIDFromFileName(fileName),
		Title: title,
		File:  fileName,
		Subs:  subPages,
	}, nil
}

func DefaultParser() *parser.Parser {
	return parser.New()
}
