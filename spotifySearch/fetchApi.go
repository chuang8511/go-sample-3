package spotifySearch


import (
	"io"
	"net/http"
	"fmt"
	"sync"
)


func FetchApi(requestUrl string, token string, wg *sync.WaitGroup, ch chan []byte) ([]byte, error) {
	
	// concurrency
	if wg != nil {
		defer wg.Done()
	}
	
	fmt.Println("Fetch Api")
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

	// concurrency
	if ch != nil {
		ch <- responseBody
	}
	fmt.Println("Get Api Response")
	return responseBody, nil
	
}