package controllers

import (
	"chatAppServer/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type Result struct {
	Status int `json:"status"`
	Msg string `json:"msg"`
}

func (this *UserController) Login() {
	var ob models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	obj := models.QuerUser(&ob)
	println(obj)
	result := Result{0, "asd"}
	this.Data["json"] = &result
	this.ServeJSON()
}

