package spotifytoken

type Token struct {
	AcessToken string `json:"access_token"`
	TokenType  string `json:"token_type"`
	ExpiresIn  int    `json:"expires_in"`
}
