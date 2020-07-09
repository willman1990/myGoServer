package models

import (
	"myGoWeb/goApiServer/dto"
	"myGoWeb/goApiServer/providers"
)

type CheckLoginModel struct {
	loginProvider *providers.LoginProvider
}

func (this *CheckLoginModel) CheckUserLogin(login *dto.Users) (*dto.LoginToken, int, error) {
	return this.loginProvider.MakeCheckForLogin(login)
}

func NewCheckLoginModel(providers *providers.Providers) *CheckLoginModel {
	return &CheckLoginModel{loginProvider: providers.LoginProvider}
}
