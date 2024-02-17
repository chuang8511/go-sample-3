package spotifyFilter

import (
	"../spotifySearch"
	"strings"
	"reflect"
)

func Filter(data spotifySearch.Episodes, key string, value string) ([]spotifySearch.Episode, error) {
	var filteredEpisodes []spotifySearch.Episode

	for _, episode := range data.Items {
		fieldValue := reflect.ValueOf(episode).FieldByName(key)
		if fieldValue.IsValid() && fieldValue.Kind() == reflect.String && strings.Contains(fieldValue.String(), value) {
            filteredEpisodes = append(filteredEpisodes, episode)
        }
	}
	
	return filteredEpisodes, nil

}