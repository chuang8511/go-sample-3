package main


import (
	"./spotifySearch"
	// "./spotifytoken"
	"fmt"
)

func main() {
	
	responseJson := spotifySearch.SearchArtist("13bJ4DAZH1QLc1fOmlZI24")
	
	fmt.Println(responseJson)
	// token, err := spotifytoken.GetToken()
	// if err != nil {
	// 	fmt.Println("errorrrrrr!!!")
	// }
	// fmt.Println(token.AcessToken)
	
}