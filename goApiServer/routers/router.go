package routers

import (
	"myGoWeb/goApiServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/will/:pa", &controllers.MainController{})
}
