package spotifySearch

import (
	"../spotifytoken"
	"encoding/json"
	"fmt"
)


func SearchArtist(id string) (*ArtistResponse, error) {

	var token string
	token, _ = spotifytoken.GetCacheToken()

	if token == "" {
		fmt.Println("Get token from Spotify API")
		tokenStruct, err := spotifytoken.GetToken()
		if err != nil {
			return nil, fmt.Errorf("cannot get token from Spotify API: %w", err)
		}
		token = tokenStruct.AcessToken
	}	

	artistId := id
	spotifyDomainUrl := "https://api.spotify.com/v1/"
	category := "shows/"

	requestUrl := spotifyDomainUrl + category + artistId

	responseBody, _ := FetchApi(requestUrl, token)

	var responseJson ArtistResponse

	err := json.Unmarshal(responseBody, &responseJson)
	
	if err != nil {
		return nil, fmt.Errorf("cannot read response: %w", err)
	}

	for responseJson.Episodes.Next != "" {
		responseBody, _ := FetchApi(responseJson.Episodes.Next, token)
		var episodes Episodes

		json.Unmarshal(responseBody, &episodes)
		responseJson.Episodes.Next = episodes.Next
		responseJson.Episodes.Items = append(responseJson.Episodes.Items, episodes.Items...)
	}

	return &responseJson, nil
}