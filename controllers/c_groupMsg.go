package controllers

import (
	"chatAppServer/models"
	"encoding/json"
)

/*GroupMsgController 组消息控制器 */
type GroupMsgController struct{
	UserController
}

/*groupMsgResult 返回结果 */
type groupMsgResult struct {
	Status int             `json:"status"`
	Msg    string          `json:"msg"`
	Data   []models.SysGroupMsg `json:"data" description:"群组消息"`
}

/*AddGroupMsg 添加新的组消息接口 */
// @Title 添加新的组消息接口
// @Description 添加新的组消息接口
// @Param data body models.SysGroupMsg true "请求参数"
// @Success 200 {object} controllers.c_groupMsg.groupMsgResult
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /	 [Post]
func (c *GroupMsgController) AddGroupMsg() {
	var ob models.SysGroupMsg
	result := groupMsgResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	res := models.AddGroupMsg(&ob)
	if res == 0 {
		result.Status = 0
		result.Msg = "群消息发送失败"
		c.Data["json"] = &result
		c.ServeJSON()
	}
	result.Status = 1
	result.Msg = "群消息发送成功"
	c.Data["json"] = &result
	c.ServeJSON()
}

/*GetGroupMsgList 获取用户消息*/
// @Title 获取用户消息
// @Description 获取用户消息
// @Param data body models.SysMsg true "请求参数"
// @Success 200 {object} controllers.c_groupMsg.groupMsgResult
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /getGroupMsgList [Post]
func (c *GroupMsgController) GetGroupMsgList() {
	var ob models.SysGroupMsg
	result := groupMsgResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	resData := models.GetGroupMsgList(&ob)
	result.Status = 1
	result.Msg = "获取群消息成功"
	// 取出用户信息
	result.Data = resData
	c.Data["json"] = &result
	c.ServeJSON()
}
