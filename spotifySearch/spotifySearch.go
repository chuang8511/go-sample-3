package spotifySearch

import (
	"../spotifytoken"
	"io"
	"net/http"
	"encoding/json"
	"fmt"
)


func SearchArtist(id string) (*Response, error) {

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

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot build request: %w", err)
	}
	
	req.Header.Set("Authorization", "Bearer " + token)

	client := &http.Client{}
	
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("cannot send request: %w", err)
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("cannot read response: %w", err)
	}

	var responseJson Response

	erro := json.Unmarshal(responseBody, &responseJson)
	
	if erro != nil {
		return nil, fmt.Errorf("cannot read response: %w", err)
	}

	return &responseJson, nil
}