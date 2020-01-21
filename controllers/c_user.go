package controllers

import (
	"fmt"
	"chatAppServer/models"
	"encoding/json"
)


/*UserController 用户控制器 */
type UserController struct {
	UtilsController
	TokenController
}

/*Code 验证码 */
type Code struct {
	Code string `json:"code" required:"true" description:"短信验证码"`
}

/*Result 返回结果 */
type Result struct {
	Status int            `json:"status"`
	Msg    string         `json:"msg"`
}

/*LoginResult 返回结果 */
type LoginResult struct {
	Status int            `json:"status"`
	Msg    string         `json:"msg"`
	Data   models.SysUser `json:"data"`
	T   string `json:"t"`
}

/*Login 登陆注册接口 */
// @Title 登陆
// @Description 用户登陆接口
// @Param data body models.SysUser true "请求参数"
// @Success 200 {object} controllers.c_user.Result
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /login [Post]
func (c *UserController) Login() {
	var ob models.SysUser
	result := Result{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	resData := models.QuerUser(&ob)
	// 保存用户信息到session
	c.SetSession("userInfo", resData[0])
	if len(resData) == 0 {
		id := models.AddUser(&ob)
		if id == 0 {
			result.Status = 0
			result.Msg = "短信验证码发送失败"
			c.Data["json"] = &result
			c.ServeJSON()
		}
	}
	// 发送短信验证码
	// flag := UtilsController.SendMessage(ob.Phone)
	flag := c.SendMessage(ob.Phone)
	fmt.Println(flag)
	if flag {
		result.Status = 1
		result.Msg = "短信验证码发送成功"
	} else {
		result.Status = 0
		result.Msg = "短信验证码发送失败"
	}
	c.Data["json"] = &result
	c.ServeJSON()
}

/*CheckCode 验证登陆验证码 */
// @Title 登陆验证码
// @Description 验证登陆验证码接口
// @Param data body controllers.c_user.Code true "请求参数"
// @Success 200 {object} controllers.c_user.Result
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /checkCode [Post]
func (c *UserController) CheckCode() {
	var ob Code
	result := LoginResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	// 发送短信验证码
	code := c.GetSession("code")
	userInfo := (c.GetSession("userInfo")).(models.SysUser)
	if ob.Code == code {
		token := c.CreateToken(userInfo.Phone)
		result.Status = 1
		result.Msg = "登陆成功"
		result.T = token
		// 取出用户信息
		result.Data = userInfo
		
	} else {
		result.Status = 0
		result.Msg = "短信验证码错误"
	}
	c.Data["json"] = &result
	c.ServeJSON()
}
