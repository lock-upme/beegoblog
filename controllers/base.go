package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	isLogin bool
}

func (this *BaseController) Prepare() {
	userLogin := this.GetSession("userLogin")
	if userLogin == nil {
		this.isLogin = false
	} else {
		this.isLogin = true
	}
	this.Data["isLogin"] = this.isLogin
}

func (this *BaseController) Go404() {
	this.TplName = "404.tpl"
	return
}
