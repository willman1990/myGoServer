package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	pa := c.Ctx.Input.Param(":pa")
	userName := c.Ctx.Input.Query("userName")
	passWord := c.Ctx.Input.Query("password")
	fmt.Printf("paramsList is ===>  pa: %s name: %s  pwd: %s" , string(pa), string(userName), string(passWord))


	fmt.Println()
	fmt.Println("======================================================================================")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
