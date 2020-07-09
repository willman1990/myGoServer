package models

import "myGoWeb/goApiServer/providers"


var LoginModel *CheckLoginModel

func NewModels(providers *providers.Providers) {
	LoginModel = NewCheckLoginModel(providers)
}
