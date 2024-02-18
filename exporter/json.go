package exporter

import (
	"../spotifySelector"
	"os"
	"encoding/json"
)

type JsonExporter struct{}

func (e JsonExporter) Export(data interface{}, filename string) error {

	response, _ := data.(*spotifySelector.Response)
	exportFilename := filename + ".json"
	file, _ := os.Create(exportFilename)

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(response)

	return nil
}