package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"myGoWeb/goApiServer/app"
	"myGoWeb/goApiServer/dto"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Post() {
	var loginInfo dto.Users
	reqBody := this.Ctx.Input.RequestBody
	err := json.Unmarshal(reqBody, &loginInfo)
	if err != nil {
		fmt.Println("Unmarshal loginInfo failed, errMsg : " + err.Error())
	}
	token, _, err := app.CheckLoginInfo(&loginInfo)
	if err != nil {
		this.CustomAbort(403, err.Error())
	}
	this.Data["json"] = token
	this.ServeJSON()

}