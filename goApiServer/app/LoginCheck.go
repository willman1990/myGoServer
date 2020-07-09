package app

import (
	"myGoWeb/goApiServer/dto"
	"myGoWeb/goApiServer/models"
)

type LoginCheck struct {

}

func CheckLoginInfo(login *dto.Users) (*dto.LoginToken, int, error) {

	resp, code, err := models.LoginModel.CheckUserLogin(login)
	if err != nil {
		return nil, code, err
	}
	return resp, 200, nil
}
