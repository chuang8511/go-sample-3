package spotifytoken

import(
	"bytes"
	"net/http"
	"io"
	"encoding/json"
	"time"
	"github.com/redis/go-redis"
	"context"
)


func expiredDateTime(expiresIn int) time.Time {
	currentDateTime := time.Now()
	expireDateTime := currentDateTime.Add(time.Second * time.Duration(expiresIn))
	return expireDateTime
}


func GetToken() (*Token, error) {
	
	grantType := "client_credentials"
	clientId := "" // remember to mock it
	clientSecret := "" // remember to mock it
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

	var ctx = context.Background()
	
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
        Password: "",
        DB:       0,
	})

	expiredDateTime := expiredDateTime(token.ExpiresIn)
	tokenSession := make(map[string]interface{})
	tokenSession["token"] = token.AcessToken
	tokenSession["expiredDateTime"] = expiredDateTime
	rdb.HSet(ctx, "tokenSession", tokenSession)

	return &token, erro
	
}