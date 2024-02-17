package spotifySearch


import (
	"io"
	"net/http"
	"fmt"
)


func FetchApi(requestUrl string, token string) ([]byte, error) {
	
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
	return responseBody, nil
}