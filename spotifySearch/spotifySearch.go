package spotifySearch

import (
	"../spotifytoken"
	"io"
	"net/http"
	"encoding/json"
)


func SearchArtist(id string) (map[string]interface{}) {

	artistId := id
	spotifyDomainUrl := "https://api.spotify.com/v1/"
	category := "shows/"

	requestUrl := spotifyDomainUrl + category + artistId

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return map[string]interface{}{ "message": "cannot build request" }
	}

	token, err := spotifytoken.GetToken()

	if err != nil {
		return map[string]interface{}{ "message": "cannot get token" }
	}

	accessTokenString := token.AcessToken
	
	
	req.Header.Set("Authorization", "Bearer " + accessTokenString)

	client := &http.Client{}
	
	resp, err := client.Do(req)

	if err != nil {
		return map[string]interface{}{ "message": "cannot send request" }
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return map[string]interface{}{ "message": "cannot read response" }
	}

	var responseJson map[string]interface{}

	erro := json.Unmarshal(responseBody, &responseJson)
	
	if erro != nil {
		return map[string]interface{}{ "message": "cannot read response" }
	}

	return responseJson
}