package navigator

import "fmt"

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
