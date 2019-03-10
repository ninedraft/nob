package navigator

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"github.com/stretchr/testify/assert"
)

func TestParsePage(test *testing.T) {
	var pageWithSubPages = "testdata/pageWithSubPages.md"
	var data, errReadFile = ioutil.ReadFile(pageWithSubPages)
	if errReadFile != nil {
		test.Fatal(errReadFile)
	}
	var page, errParsePage = ParsePage(pageWithSubPages, bytes.NewReader(data))
	if errParsePage != nil {
		test.Fatal(errParsePage)
	}
	assert.Equal(test, Page{}, page)
}

func TestParser(test *testing.T) {
	var pageWithSubPages = "testdata/pageWithSubPages.md"
	var data, errReadFile = ioutil.ReadFile(pageWithSubPages)
	if errReadFile != nil {
		test.Fatal(errReadFile)
	}
	var walker = ast.NodeVisitorFunc(func(node ast.Node, entering bool) ast.WalkStatus {
		switch node := node.(type) {
		case *ast.Heading:
			var buf = &bytes.Buffer{}
			var text = node.GetChildren()[0].(*ast.Text)
			ast.Print(buf, text)
			test.Log(buf)
			test.Logf("MANUAL: [%d] %q", node.Level, text.Literal)
		case *ast.Link:
			test.Logf("%q", node.Destination)
		}
		return ast.GoToNext
	})
	var AST = parser.New().Parse(data)
	ast.Walk(AST, walker)
}
