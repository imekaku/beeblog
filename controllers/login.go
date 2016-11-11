package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	isExit := this.Input().Get("exit") == "true"
	if isExit {
		this.Ctx.SetCookie("username", "", -1, "/")
		this.Ctx.SetCookie("password", "", -1, "/")
	}
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	autologin := this.Input().Get("autologin") == "on"

	if username == beego.AppConfig.String("username") && password == beego.AppConfig.String("password") {
		maxAge := 0
		if autologin {
			maxAge = 1<<32 - 1
		}
		this.Ctx.SetCookie("username", username, maxAge, "/")
		this.Ctx.SetCookie("password", password, maxAge, "/")
	}
	this.Redirect("/", 301)
	return
}

func checkAccout(ctx *context.Context) bool {
	ckUsername := ctx.GetCookie("username")
	ckPassword := ctx.GetCookie("password")
	if ckUsername == beego.AppConfig.String("username") && ckPassword == beego.AppConfig.String("password") {
		return true
	} else {
		return false
	}
}
