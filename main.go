package main

import (
	"fmt"
	"./spotifyFilter"
	"./spotifySearch"
	"./spotifySelector"
)

func main() {
	
	responseJson, _ := spotifySearch.SearchArtist("13bJ4DAZH1QLc1fOmlZI24")	

	fmt.Println(responseJson.Episodes.Limit)
	fmt.Println(responseJson.Episodes.Next)
	fmt.Println(responseJson.Episodes.Offset)

	res, err := spotifyFilter.Filter(responseJson.Episodes, "Name", "新資料夾")
	
	if err != nil {
		fmt.Println("err!!!!!!!!")
	}

	display, err := spotifySelector.Select(res)

	if err != nil {
		fmt.Println("err!!!!!!!!")
	}
	
	fmt.Println(display)
}