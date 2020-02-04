package controllers

import (
	"chatAppServer/models"
	"encoding/json"
)

/*MsgController 消息控制器 */
type MsgController struct{
	UserController
}

/*MsgResult 返回结果 */
type MsgResult struct {
	Status int             `json:"status"`
	Msg    string          `json:"msg"`
	Data   []models.SysMsg `json:"data" description:"用户消息"`
}

/*AddMsg 添加新消息接口 */
// @Title 添加新消息接口
// @Description 添加新消息接口
// @Param data body models.SysMsg true "请求参数"
// @Success 200 {object} controllers.c_msg.Result
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /addMsg [Post]
func (c *MsgController) AddMsg() {
	var ob models.SysMsg
	result := MsgResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	res := models.AddMsg(&ob)
	if res == 0 {
		result.Status = 0
		result.Msg = "消息发送失败"
		c.Data["json"] = &result
		c.ServeJSON()
	}
	result.Status = 1
	result.Msg = "消息发送成功"
	c.Data["json"] = &result
	c.ServeJSON()
}

/*GetMsgList 获取用户消息*/
// @Title 获取用户消息
// @Description 获取用户消息
// @Param data body models.SysMsg true "请求参数"
// @Success 200 {object} controllers.c_msg.Result
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /getMsgList [Post]
func (c *MsgController) GetMsgList() {
	var ob models.SysMsg
	result := MsgResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	resData := models.GetMsgList(&ob)
	result.Status = 1
	result.Msg = "获取消息成功"
	// 取出用户信息
	result.Data = resData
	c.Data["json"] = &result
	c.ServeJSON()
}
