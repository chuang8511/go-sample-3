package spotifySearch

import (
	"../spotifytoken"
	"encoding/json"
	"fmt"
	"sync"
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
	fmt.Println("token: ", token)

	artistId := id
	spotifyDomainUrl := "https://api.spotify.com/v1/"
	category := "shows/"

	requestUrl := spotifyDomainUrl + category + artistId

	responseBody, _ := FetchApi(requestUrl, token, nil, nil)

	var responseJson ArtistResponse

	err := json.Unmarshal(responseBody, &responseJson)
	
	if err != nil {
		return nil, fmt.Errorf("cannot read response: %w", err)
	}

	totalCount := responseJson.Episodes.Total
	onePageCount := responseJson.Episodes.Limit

	loopCount := totalCount / onePageCount - 1

	var waitGroup sync.WaitGroup
	ch := make(chan []byte, loopCount)

	for i := 1; i < loopCount; i++ {
		waitGroup.Add(1)
		go FetchApi(responseJson.Episodes.Next, token, &waitGroup, ch)
		
	}

	go func() {
        waitGroup.Wait()
        close(ch)
    }()
	
	for responseBody := range ch {
        var episodes Episodes

		if err := json.Unmarshal(responseBody, &episodes); err != nil {
            fmt.Println("Error unmarshalling response:", err)
            continue
        }
        responseJson.Episodes.Items = append(responseJson.Episodes.Items, episodes.Items...)
    }

	return &responseJson, nil
}