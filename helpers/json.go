package helpers

import (
	"../models"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
)

type HelpersJson struct {

}

func (u *HelpersJson)readFileJson(location string)(outputStruct interface{})  {
	file, err := ioutil.ReadFile(location)

	if err != nil{
		fmt.Println("read 1",err.Error())
	}

	var data interface{}

	err = json.Unmarshal([]byte(file), &data)

	if err != nil{
		fmt.Println("read 2",err.Error())
	}

	return data
}

func (u *HelpersJson)WriteFileJsonAuthLogin(dataAuth models.Auth)  {

	data := u.readFileJson("./documents/auth.json")
	fmt.Println(data)
	c := models.JsonAuth{}
	err := mapstructure.Decode(data, &c)

	if err != nil{
		fmt.Println("write 1",err.Error())
	}

	fmt.Println(c)
	auth := []models.Auth{}

	for i := range c.Auth_customer {
		if c.Auth_customer[i].User_name != dataAuth.User_name{
			auth = append(auth, c.Auth_customer[i])
		}
	}

	auth = append(auth, dataAuth)

	jsonAuth := models.JsonAuth{}

	jsonAuth.Auth_customer = auth


	file, err := json.Marshal(jsonAuth)

	if err != nil{
		fmt.Println("write 2",err.Error())
	}

	err = ioutil.WriteFile("./documents/auth.json", file, 0644)

	if err != nil{
		fmt.Println("write 3",err.Error())
	}
}
func (u *HelpersJson)WriteFileJsonAuthLogout(dataAuth models.Auth)  {

	data := u.readFileJson("./documents/auth.json")
	fmt.Println(data)
	c := models.JsonAuth{}
	err := mapstructure.Decode(data, &c)

	if err != nil{
		fmt.Println("write 1",err.Error())
	}

	fmt.Println(c)
	auth := []models.Auth{}

	for i := range c.Auth_customer {
		if c.Auth_customer[i].User_name != dataAuth.User_name{
			auth = append(auth, c.Auth_customer[i])
		}
	}

	jsonAuth := models.JsonAuth{}

	jsonAuth.Auth_customer = auth


	file, err := json.Marshal(jsonAuth)

	if err != nil{
		fmt.Println("write 2",err.Error())
	}

	err = ioutil.WriteFile("./documents/auth.json", file, 0644)

	if err != nil{
		fmt.Println("write 3",err.Error())
	}
}
func (u *HelpersJson)WriteFileJsonRiwayatPembayaran(dataRiwayat models.Riwayat)  {

	data := u.readFileJson("./documents/riwayat.json")
	fmt.Println(data)
	c := models.JsonRiwayat{}
	err := mapstructure.Decode(data, &c)

	if err != nil{
		fmt.Println("write 1",err.Error())
	}

	fmt.Println(c)
	riw := []models.Riwayat{}

	for i := range c.Riwayat {
		riw = append(riw, c.Riwayat[i])
	}

	riw = append(riw, dataRiwayat)

	jsonRiwayat := models.JsonRiwayat{}

	jsonRiwayat.Riwayat = riw


	file, err := json.Marshal(jsonRiwayat)

	if err != nil{
		fmt.Println("write 2",err.Error())
	}

	err = ioutil.WriteFile("./documents/riwayat.json", file, 0644)

	if err != nil{
		fmt.Println("write 3",err.Error())
	}
}


func (u *HelpersJson)GetCustomer() []models.Customer {

	data := u.readFileJson("./documents/customer.json")
	c := models.JsonCustomer{}
	err := mapstructure.Decode(data, &c)

	if err != nil{
		fmt.Println(err.Error())
	}

	customer := []models.Customer{}

	for i := range c.Customer {
		customer = append(customer, c.Customer[i])
	}

	return customer
}
func (u *HelpersJson)GetMerchant() []models.Merchant {

	data := u.readFileJson("./documents/merchant.json")
	c := models.JsonMerchant{}
	err := mapstructure.Decode(data, &c)

	if err != nil{
		fmt.Println(err.Error())
	}

	merchant := []models.Merchant{}

	for i := range c.Merchant {
		merchant = append(merchant, c.Merchant[i])
	}

	return merchant
}
func (u *HelpersJson)GetRiwayat() []models.Riwayat {

	data := u.readFileJson("./documents/riwayat.json")
	c := models.JsonRiwayat{}
	err := mapstructure.Decode(data, &c)

	if err != nil{
		fmt.Println(err.Error())
	}

	riw := []models.Riwayat{}

	for i := range c.Riwayat {
		riw = append(riw, c.Riwayat[i])
	}

	return riw
}
func (u *HelpersJson)GetAuth() []models.Auth {

	data := u.readFileJson("./documents/auth.json")
	c := models.JsonAuth{}
	err := mapstructure.Decode(data, &c)

	if err != nil{
		fmt.Println(err.Error())
	}

	auth := []models.Auth{}

	for i := range c.Auth_customer {
		auth = append(auth, c.Auth_customer[i])
	}

	return auth
}