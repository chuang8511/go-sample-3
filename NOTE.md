# About Go
go mod tidy -> like bundle install

# About RESTful API
Ref:
- https://www.youtube.com/watch?v=d_L64KT3SFM&ab_channel=LaithAcademy

bug
- main.go:6:2: no required module provides package github.com/gin-gonic/gin: go.mod file not found in current directory or any parent directory; see 'go help modules'

Solution
add this line in your ~/.zprofile
export GO111MODULE=off

go get -u github.com/gin-gonic/gin

go run main.go



Testing command
curl -d '{"id": "4", "title": "HTTP", "completed": false}' -X POST http://localhost:9090/todos

curl "http://localhost:9090/todos"

curl -X PATCH "http://localhost:9090/todos/1"

# About Spotify API
curl -X POST "https://accounts.spotify.com/api/token" \
     -H "Content-Type: application/x-www-form-urlencoded" \
     -d "grant_type=client_credentials&client_id=xx&client_secret=xx"

https://accounts.spotify.com/authorize?client_id=client_id&response_type=token&redirect_uri=http://localhost:8080/authorize&scope=user-read-private&user-read-email


{"access_token":"BQAAX5jipGBDnyTkzYpRgP8Mq62jnrrvyridItCeRNXTJsMF5Eo5TUNy2is1viv7dO8VxnxnX_cCwPxuNi3ZR0W8nnRQtVHQSB2hbkUOyQuwateJ3Oo","token_type":"Bearer","expires_in":3600}

curl "https://api.spotify.com/v1/artists/4Z8W4fKeB5YxbusRsdQVPb" \
     -H "Authorization: Bearer  BQAAX5jipGBDnyTkzYpRgP8Mq62jnrrvyridItCeRNXTJsMF5Eo5TUNy2is1viv7dO8VxnxnX_cCwPxuNi3ZR0W8nnRQtVHQSB2hbkUOyQuwateJ3Oo"

{
  "external_urls" : {
    "spotify" : "https://open.spotify.com/artist/4Z8W4fKeB5YxbusRsdQVPb"
  },
  "followers" : {
    "href" : null,
    "total" : 9340749
  },
  "genres" : [ "alternative rock", "art rock", "melancholia", "oxford indie", "permanent wave", "rock" ],
  "href" : "https://api.spotify.com/v1/artists/4Z8W4fKeB5YxbusRsdQVPb",
  "id" : "4Z8W4fKeB5YxbusRsdQVPb",
  "images" : [ {
    "height" : 640,
    "url" : "https://i.scdn.co/image/ab6761610000e5eba03696716c9ee605006047fd",
    "width" : 640
  }, {
    "height" : 320,
    "url" : "https://i.scdn.co/image/ab67616100005174a03696716c9ee605006047fd",
    "width" : 320
  }, {
    "height" : 160,
    "url" : "https://i.scdn.co/image/ab6761610000f178a03696716c9ee605006047fd",
    "width" : 160
  } ],
  "name" : "Radiohead",
  "popularity" : 81,
  "type" : "artist",
  "uri" : "spotify:artist:4Z8W4fKeB5YxbusRsdQVPb"
}

curl --request GET \
  --url 'https://api.spotify.com/v1/search?q=remaster%2520track%3ADoxy%2520artist%3AMiles%2520Davis&type=artist&market=%E5%91%B1%E5%90%89' \
  --header 'Authorization: Bearer BQAAX5jipGBDnyTkzYpRgP8Mq62jnrrvyridItCeRNXTJsMF5Eo5TUNy2is1viv7dO8VxnxnX_cCwPxuNi3ZR0W8nnRQtVHQSB2hbkUOyQuwateJ3Oo'


curl --request GET \
  --url https://api.spotify.com/v1/shows/13bJ4DAZH1QLc1fOmlZI24 \
  --header 'Authorization: Bearer BQBoVSZIkR7hB8XFfJ2Uy0HIDaALO06HCGbC1cXLXPQ2CcUoL__7j9RhWq3eNisXxR2ve01Zs5xUByq535nrVqKNAHQG4Gib07vjk0z0eaYFHk_382o'



curl --request GET \
  --url https://api.spotify.com/v1/me/playlists \
  --header 'Authorization: Bearer BQBoVSZIkR7hB8XFfJ2Uy0HIDaALO06HCGbC1cXLXPQ2CcUoL__7j9RhWq3eNisXxR2ve01Zs5xUByq535nrVqKNAHQG4Gib07vjk0z0eaYFHk_382o'

  