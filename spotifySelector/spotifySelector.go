package spotifySelector

import (
	"../spotifySearch"
)

func Select(data []spotifySearch.Episode) (*Response, error) {

	var res Response

	for _, episode := range data {
		var playItem PlayItem
		id := episode.ID
		playLink := "https://open.spotify.com/episode/" + id
		playItem.Name = episode.Name
		playItem.PlayLink = playLink
		res.PlayItems = append(res.PlayItems, playItem)
	}


	return &res, nil

}
