package spotifytoken

import(
	"bytes"
	"net/http"
	"io"
	"encoding/json"
)

type Token struct {
	AcessToken string `json:"access_token"`
	TokenType  string `json:"token_type"`
	ExpiresIn  int    `json:"expires_in"`
}

func GetToken() (*Token, error) {
	
	grantType := "client_credentials"
	clientId := "mockid" // remember to mock it
	clientSecret := "mock secret" // remember to mock it
	bodyString := "grant_type=" + grantType + "&client_id=" + clientId + "&client_secret=" + clientSecret
	url := "https://accounts.spotify.com/api/token"

    requestBody := []byte(bodyString)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
    resp, err := client.Do(req)
    
	if err != nil {
		return nil, err
    }
	
    defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
    
	if err != nil {
        return nil, err
    }
	
	var token Token

	erro := json.Unmarshal(responseBody, &token)
	
	if erro != nil {
		return nil, err
	}

	return &token, erro
	
}