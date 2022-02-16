package controllers

import (
	"../middleware"
	"../models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

import (
	"../helpers"
)

type Customer struct {
	helper helpers.HelpersJson
}

func (u *Customer)Pembayaran(c *gin.Context)  {

	res := models.Response{}

	username := c.PostForm("username")
	id_merchant := c.PostForm("id_merchant")
	saldo := c.PostForm("saldo")

	if username == "" && id_merchant == "" && saldo == ""{

		res.ApiMessage = "username dan password wajib diisi"
		c.JSON(400,res)
		return

	}

	outCustomer := u.helper.GetCustomer()
	dataCustomer := models.Customer{}

	var akses bool
	for i := range outCustomer {

		if outCustomer[i].User_name == username{
			dataCustomer = outCustomer[i]
			akses = true
			break
		}
	}

	if !akses{
		res.ApiMessage = "Pengguna tidak terdaftar , tidak dapat melakukan pembayaran"
		c.JSON(400,res)
		return
	}

	outMerchant := u.helper.GetMerchant()

	for i := range outMerchant {

		if outMerchant[i].Id == id_merchant{
			akses = true
			break
		}
	}

	if !akses{
		res.ApiMessage = "Merchant tidak ditemukan"
		c.JSON(400,res)
		return
	}

	saldoInt , _ := strconv.Atoi(saldo)

	sisaSaldo := dataCustomer.Saldo - saldoInt

	if sisaSaldo <= 0 {
		res.ApiMessage = "Saldo anda tidak cukup"
		c.JSON(400,res)
		return
	}

	u.helper.WriteFileJsonRiwayatPembayaran(models.Riwayat{
		Nominal_pembayaran: saldoInt,
	User_customer: username,
	User_merchant: id_merchant,
	Tanggal_pembayaran: time.Now().String()})

	outRiwayat := u.helper.GetRiwayat()
	dataRiwayat := []models.Riwayat{}

	for i := range outRiwayat {
		if outRiwayat[i].User_customer == username{
			dataRiwayat = append(dataRiwayat, outRiwayat[i])
		}
	}

	res.ApiMessage = "Pembayaran berhasil, berikut histori pembayaran anda"
	res.Data = dataRiwayat

	c.JSON(200,res)

}
func (u *Customer)Login(c *gin.Context)  {

	res := models.Response{}

	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" && password == ""{
		res.ApiMessage = "username dan password wajib diisi"
		c.JSON(400,res)
	}

	outCustomer := u.helper.GetCustomer()

	var akses bool
	for i := range outCustomer {

		if outCustomer[i].User_name == username{
			if outCustomer[i].Password == password{
				mw := middleware.MiddleWare{}

				token , err := mw.CreateAuth(outCustomer[i].User_name,"customer")
				if err != nil{
					res.ApiMessage = "Terdapat kesalahan silahkan ulangi login kembali"
					akses = false
				}else{

					tokenRes := models.Token{
						Token: token,
					}
					res.ApiMessage = "Login Berhasil"
					res.Data = tokenRes
					akses = true
				}

			}else{
				res.ApiMessage = "Password anda salah"
				akses = false
			}
			break
		}else{
			res.ApiMessage = "Akun anda tidak ditemukan"
		}
	}

	if akses{
		c.JSON(200,res)
	}else{
		c.JSON(400,res)
	}

}
func (u *Customer)Logout(c *gin.Context)  {

	res := models.Response{}

	username := c.PostForm("username")
	token := c.Request.Header.Get("Authorization")

	outAuth := u.helper.GetAuth()

	var akses bool
	for i := range outAuth {

		if outAuth[i].User_name == username{
			if outAuth[i].Token == token{

				u.helper.WriteFileJsonAuthLogout(models.Auth{User_name: username,Token: token})
				res.ApiMessage = "Logout Berhasil"
				akses = true

			}else{
				res.ApiMessage = "Data tidak cocok , gagal logout"
				akses = false
			}
			break
		}else{
			res.ApiMessage = "Data tidak ditemukan , gagal lout"
		}
	}

	if akses{
		c.JSON(200,res)
	}else{
		c.JSON(400,res)
	}

}


