package spotifySearch

import (
	"../spotifytoken"
	"io"
	"net/http"
	"encoding/json"
	"github.com/redis/go-redis"
	"context"
	"fmt"
	"time"
)

func expiredDateTime(expiresIn int) time.Time {
	currentDateTime := time.Now()
	expireDateTime := currentDateTime.Add(time.Second * time.Duration(expiresIn))
	return expireDateTime
}


func SearchArtist(id string) (map[string]interface{}, error) {

	var token string
	var ctx = context.Background()
	
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
        Password: "",
        DB:       0,
	})

	expiredDateTimeString, errDateTime := rdb.HGet(ctx, "tokenSession", "expiredDateTime").Result()

	if errDateTime != redis.Nil {
		expiredDateTime, _ := time.Parse(time.RFC3339, expiredDateTimeString)
		currentDateTime := time.Now()
		if expiredDateTime.After(currentDateTime) {
			fmt.Println("Get token from Cache")
			token, _ = rdb.HGet(ctx, "tokenSession", "token").Result()
		}
	}

	if token == "" {
		fmt.Println("Get token from Spotify API")
		token, err := spotifytoken.GetToken()
		if err != nil {
			return nil, fmt.Errorf("cannot get token from Spotify API: %w", err)
		}
		expiredDateTime := expiredDateTime(token.ExpiresIn)
		tokenSession := make(map[string]interface{})
		tokenSession["token"] = token.AcessToken
		tokenSession["expiredDateTime"] = expiredDateTime
		rdb.HSet(ctx, "tokenSession", tokenSession)
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

	var responseJson map[string]interface{}

	erro := json.Unmarshal(responseBody, &responseJson)
	
	if erro != nil {
		return nil, fmt.Errorf("cannot read response: %w", err)
	}

	return responseJson, nil
}