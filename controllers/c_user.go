package controllers

import (
	"fmt"
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

// @Title 登陆
// @Description 用户登陆接口
// @Param phone formData string false "手机号"
// @Param id formData string false "用户id"
// @Param nickName formData string false "用户名称"
// @Success 200 {object} controllers.c_user.Result
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /user/login [Post]
func (this *UserController) Post() {
	var ob models.SysUser
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	flag := models.QuerUser(&ob)
	fmt.Print(flag)
	result := Result{0, "asd"}
	this.Data["json"] = &result
	this.ServeJSON()
}

