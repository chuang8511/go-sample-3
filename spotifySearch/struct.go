package spotifySearch

type Episode struct {
	AudioPreviewURL       string   `json:"audio_preview_url"`
	Description           string   `json:"description"`
	// HTMLDescription       string   `json:"html_description"`
	DurationMS            int      `json:"duration_ms"`
	Explicit              bool     `json:"explicit"`
	ExternalURLs          struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href           string `json:"href"`
	ID             string `json:"id"`
	// Images         []struct {
	// 	URL    string `json:"url"`
	// 	Height int    `json:"height"`
	// 	Width  int    `json:"width"`
	// } `json:"images"`
	IsExternallyHosted bool     `json:"is_externally_hosted"`
	IsPlayable         bool     `json:"is_playable"`
	Language           string   `json:"language"`
	Languages          []string `json:"languages"`
	Name               string   `json:"name"`
	ReleaseDate        string   `json:"release_date"`
	ReleaseDatePrecision string `json:"release_date_precision"`
	ResumePoint        struct {
		FullyPlayed         bool `json:"fully_played"`
		ResumePositionMS    int  `json:"resume_position_ms"`
	} `json:"resume_point"`
	Type        string `json:"type"`
	URI         string `json:"uri"`
	// Restrictions struct {
	// 	Reason string `json:"reason"`
	// } `json:"restrictions"`
}

type Episodes struct {
	// Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	// Previous string `json:"previous"`
	// Total    int    `json:"total"`
	Items    []Episode `json:"items"`
}

type ArtistResponse struct {
	// AvailableMarkets []string `json:"available_markets"`
	Episodes         Episodes `json:"episodes"`
}