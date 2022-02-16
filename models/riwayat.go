package models

type JsonRiwayat struct {
	Riwayat []Riwayat `json:"riwayat"`
}
type Riwayat struct {
	User_customer string `json:"user_customer"`
	User_merchant string `json:"user_merchant"`
	Nominal_pembayaran int `json:"nominal_pembayaran"`
	Tanggal_pembayaran string `json:"tanggal_pembayaran"`
}
