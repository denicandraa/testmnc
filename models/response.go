package models

type Response struct {
	ApiMessage string `json:"api_message"`
	Data interface{} `json:"data"`
}

type Token struct {
	Token string `json:"token"`
}
