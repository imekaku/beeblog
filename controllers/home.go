package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["IsHome"] = true
	this.Data["IsLogin"] = checkAccout(this.Ctx)
	this.TplName = "home.html"
}
