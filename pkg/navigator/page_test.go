package navigator

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestParsePage(test *testing.T) {
	var pageWithSubPages = "testdata/pageWithSubPages.md"
	var data, errReadFile = ioutil.ReadFile(pageWithSubPages)
	if errReadFile != nil {
		test.Fatal(errReadFile)
	}
	var _, errParsePage = ParsePage(pageWithSubPages, bytes.NewReader(data))
	if errParsePage != nil {
		test.Fatal(errParsePage)
	}
}
