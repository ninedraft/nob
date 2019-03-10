package navigator

type Root struct {
	pages         map[ID]Page
	topLevelPages []ID
}
