package controllers

import (
	"chatAppServer/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

/*UserController 用户控制器 */
type UserController struct {
	beego.Controller
}

/*Result 返回结果 */
type Result struct {
	Status int            `json:"status"`
	Msg    string         `json:"msg"`
	Data   models.SysUser `json:"data"`
}

/*Post 登陆注册接口 */
// @Title 登陆
// @Description 用户登陆接口
// @Param data body models.SysUser true "请求参数"
// @Success 200 {object} controllers.c_user.Result
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /login [Post]
func (c *UserController) Post() {
	var ob models.SysUser
	result := Result{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	resData := models.QuerUser(&ob)
	if len(resData) > 0 {
		// 发送短信验证码
		flag := models.SendMessage(ob.Phone)
		if flag {
			result.Msg = "短信验证码发送成功"
		} else {
			result.Msg = "短信验证码发送失败"
		}
	} else {
		id := models.AddUser(&ob)
		if id > 0 {
			// 发送短信验证码
		} else {
			result.Status = 0
			result.Msg = "短信验证码发送失败"
			c.Data["json"] = &result
			c.ServeJSON()
		}
	}
	result.Status = 1
	c.Data["json"] = &result
	c.ServeJSON()
}
