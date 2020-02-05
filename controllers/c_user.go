package controllers

import (
	"chatAppServer/models"
	"encoding/json"
	"fmt"
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
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

/*LoginResult 返回结果 */
type LoginResult struct {
	Status int            `json:"status"`
	Msg    string         `json:"msg"`
	Data   models.SysUser `json:"data" description:"用户数据"`
	T      string         `json:"t" description:"token"`
}

/*SearchUserResult 返回结果 */
type SearchUserResult struct {
	Status int              `json:"status"`
	Msg    string           `json:"msg"`
	Data   []models.SysUser `json:"data" description:"用户数据"`
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
	var ob models.SearchUser
	result := Result{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	models.AddUser(&ob)
	resData := models.QuerUser(&ob)
	// 保存用户信息到session
	c.SetSession("userInfo", resData[0])
	// 发送短信验证码
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
// @Success 200 {object} controllers.c_user.LoginResult
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /checkCode [Post]
func (c *UserController) CheckCode() {
	var ob Code
	result := LoginResult{}
	errResult := Result{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	// 发送短信验证码
	code := c.GetSession("code")
	if ob.Code == code {
		// 取出用户数据
		userInfo := (c.GetSession("userInfo")).(models.SysUser)
		token := c.CreateToken(userInfo.Phone)
		result.Status = 1
		result.Msg = "登陆成功"
		result.T = token
		// 取出用户信息
		result.Data = userInfo
		c.Data["json"] = &result
	} else {
		errResult.Status = 0
		errResult.Msg = "短信验证码错误"
		c.Data["json"] = &errResult
	}
	c.ServeJSON()
}

/*SearchUser 搜索用户*/
// @Title 搜索用户
// @Description 搜索用户
// @Param data body models.SearchUser true "请求参数"
// @Success 200 {object} []models.SysUser
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /searchUser [Post]
func (c *UserController) SearchUser() {
	var ob models.SearchUser
	result := SearchUserResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	resData := models.QuerUser(&ob)
	fmt.Println(resData)
	result.Status = 1
	result.Msg = "查询用户成功"
	// 取出用户信息
	result.Data = resData
	c.Data["json"] = &result
	c.ServeJSON()
}
