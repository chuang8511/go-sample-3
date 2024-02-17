package spotifySelector

type PlayItem struct {
	Name string
	PlayLink string
}

type Response struct {
	PlayItems []PlayItem
}