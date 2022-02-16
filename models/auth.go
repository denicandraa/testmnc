package models

type JsonAuth struct {
	Auth_customer []Auth `json:"auth_customer"`
}

type Auth struct {
	User_name string `json:"user_name"`
	Token string `json:"token"`
}

type TokenClaim struct {
	Authorized      bool
	Client_Username string
	Client_Role     string
}
