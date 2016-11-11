package routers

import (
	"github.com/astaxie/beego"
	"github.com/beeblog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
}
