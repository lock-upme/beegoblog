package controllers

import (
	. "github.com/lock-upme/beegoblog/models"
	//"strconv"

	//"github.com/astaxie/beego"
)

//login
type LoginUserController struct {
	//beego.Controller
	BaseController
}

func (this *LoginUserController) Get() {
	check := this.isLogin
	if check {
		this.Redirect("/article", 302)
	} else {
		this.TplName = "login.tpl"
	}

}

func (this *LoginUserController) Post() {
	phone := this.GetString("phone")
	password := this.GetString("password")

	if "" == phone {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写手机号"}
		this.ServeJSON()
	}

	if "" == password {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写密码"}
		this.ServeJSON()
	}

	err, user := LoginUser(phone, password)

	if err == nil {
		this.SetSession("userLogin", "1")
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "贺喜你，登录成功", "user": user}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
	}
	this.ServeJSON()
}

//logout

type LogoutUserController struct {
	BaseController
}

func (this *LogoutUserController) Get() {
	this.DelSession("userLogin")
	//this.Ctx.WriteString("you have logout")
	this.Redirect("/article", 302)

}

//about me
type AboutUserController struct {
	BaseController
}

func (this *AboutUserController) Get() {

	id := 1

	pro, err := GetProfile(id)
	if err != nil {
		this.Redirect("/404.html", 302)
	}
	this.Data["pro"] = pro
	this.TplName = "about.tpl"
}
