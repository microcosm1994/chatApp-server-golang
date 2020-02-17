package controllers

import (
	"chatAppServer/models"
	"encoding/json"
)

/*VideoGroupMemberController 视频会议成员控制器 */
type VideoGroupMemberController struct {
	UserController
}

/*VideoGroupMemberResult 返回结果 */
type VideoGroupMemberResult struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   interface{} `json:"data"`
}

/*AddVideoGroupMember 新增视频会议成员 */
// @Title 新增视频会议成员
// @Description 新增视频会议成员
// @Param data body models.SysVideoGroupMember true "请求参数"
// @Success 200 {object} controllers.c_group.VideoGroupMemberResult
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /addVideoGroupMember [Post]
func (c *GroupController) AddVideoGroupMember() {
	var ob models.SysVideoGroupMember
	videoGroup := models.SysVideoGroup{}
	user := models.SysUser{}
	ob.User = &user
	ob.VideoGroup = &videoGroup
	result := VideoGroupMemberResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	res := models.AddVideoGroupMember(&ob)
	if res == 0 {
		result.Status = 0
		result.Msg = "用户已经加入视频会议"
		c.Data["json"] = &result
		c.ServeJSON()
	}
	result.Status = 1
	result.Msg = "添加会议成员成功"
	result.Data = res
	c.Data["json"] = &result
	c.ServeJSON()
}

/*GetVideoGroupUserList 获取视频会议成员*/
// @Title 获取视频会议成员
// @Description 获取视频会议成员
// @Param data body models.SysVideoGroup true "请求参数"
// @Success 200 {object} controllers.c_group.VideoGroupMemberResult
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /getVideoGroupUserList [Post]
func (c *GroupController) GetVideoGroupUserList() {
	var ob models.SysVideoGroup
	result := VideoGroupMemberResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	resData := models.GetVideoGroupUser(&ob)
	result.Status = 1
	result.Msg = "获取群成员成功"
	result.Data = resData
	c.Data["json"] = &result
	c.ServeJSON()
}

/*DelVideoGroupUser 移除视频会议成员*/
// @Title 移除视频会议成员
// @Description 移除视频会议成员
// @Param data body models.SysVideoGroupMember true "请求参数"
// @Success 200 {object} controllers.c_group.VideoGroupMemberResult
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /delVideoGroupUser [Post]
func (c *GroupController) DelVideoGroupUser() {
	var ob models.SysVideoGroupMember
	result := VideoGroupMemberResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	res := models.DelVideoGroupUser(&ob)
	if res == 0 {
		result.Status = 0
		result.Msg = "移除群成员失败"
	}
	result.Status = 1
	result.Msg = "移除群成员成功"
	c.Data["json"] = &result
	c.ServeJSON()
}
