package exporter

import (
    "encoding/csv"
    "os"
	"../spotifySelector"
)

func ExportToCSV(data *spotifySelector.Response, fileName string, headers []string) {

	file, _ := os.Create(fileName)

	defer file.Close()	
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write(headers)

	for _, item := range data.PlayItems {
        row := []string{ item.Name , item.Name }
		writer.Write(row)
    }

}