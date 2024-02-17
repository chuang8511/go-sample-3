package spotifytoken

import (
	"context"
	"github.com/redis/go-redis"
	"time"
	"fmt"
)

func GetCacheToken() (string, error) {


	var ctx = context.Background()
	var token string
	
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

	return token, nil
}