package main

import (
	// "fmt"
	"./spotifyFilter"
	"./spotifySearch"
	"./spotifySelector"
	"./exporter"
)

func main() {
	
	responseJson, _ := spotifySearch.SearchArtist("13bJ4DAZH1QLc1fOmlZI24")	

	res, _ := spotifyFilter.Filter(responseJson.Episodes, "Name", "新資料夾")
	
	display, _ := spotifySelector.Select(res)

	exporters := exporter.GetAllExporters()

	for _, e := range exporters {
		e.Export(display, "PlayList")
	}

	// fmt.Println(display)
}