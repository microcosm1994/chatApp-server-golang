package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type Result struct {
	status int
	msg string
}

func (c *UserController) GetMessageCode() {
	println(111)
	c.Data["data"] = &Result{0, "dddd"}
	c.ServeJSON()
}

func (c *UserController) CheckingCode() {
	println(222)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	c.ServeJSON()
}
