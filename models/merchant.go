package models

type JsonMerchant struct {
	Merchant []Merchant `json:"merchant.json"`
}
type Merchant struct {
	Id string `json:"id"`
	Name_merchant string `json:"name_merchant"`
}
