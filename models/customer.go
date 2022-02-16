package models

type JsonCustomer struct {
	Customer []Customer `json:"customer"`
}
type Customer struct {
	User_name string `json:"user_name"`
	Nama string `json:"nama"`
	Password string `json:"password"`
	Saldo int `json:"saldo"`
}
