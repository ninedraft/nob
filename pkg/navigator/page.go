package navigator

type ID string

type Page struct {
	ID    ID
	Title string
	File  string
	Subs  []ID
}
