package filetree

type Root struct {
	pages         map[ID]Page
	topLevelPages []ID
}

type ID string

type Page struct {
	ID    ID
	Title string
	Subs  []ID
}
