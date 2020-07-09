package dto

type Users struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

type LoginToken struct {
	Token string `json:"token"`
}