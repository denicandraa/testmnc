package main

import (
	"./controllers"
	"./middleware"
	"github.com/gin-gonic/gin"
)

type Main struct {
	router *gin.Engine
}

func (m *Main) init() error {

	m.router = gin.Default()

	return nil

}

func main() {

	//setup gin

	m := Main{}
	midd := middleware.MiddleWare{}

	if m.init() != nil {
		return
	}

	customer := controllers.Customer{}

	api := m.router.Group("/mncbank/")

	api.POST("login", customer.Login)
	api.POST("logout", customer.Logout)

	//endpoint tertentu menggunakan auth jwt
	auth := api.Group("")
	auth.Use(midd.Auth)
	{
		auth.POST("pembayaran", customer.Pembayaran)
	}

	m.router.Run(":8080")
}
