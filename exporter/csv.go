package exporter

import (
    "encoding/csv"
    "os"
	"../spotifySelector"
	"reflect"
)

type CSVExporter struct{}

func getHeaders() []string {
	var playItem spotifySelector.PlayItem

	types := reflect.TypeOf(playItem)

	headers := make([]string, types.NumField())

	for i := 0; i < types.NumField(); i++ {
		headers[i] = types.Field(i).Name
	}
	return headers
}

func (e CSVExporter) Export(data interface{}, filename string) error {
	response, _ := data.(*spotifySelector.Response)

	exportFilename := filename + ".csv"
    file, _ := os.Create(exportFilename)

	defer file.Close()	
	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := getHeaders()
	writer.Write(headers)

	for _, item := range response.PlayItems {
        row := []string{ item.Name , item.PlayLink }
		writer.Write(row)
    }

    return nil
}