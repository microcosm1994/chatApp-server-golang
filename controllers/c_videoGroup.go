package controllers

import (
	"chatAppServer/models"
	"encoding/json"
)

/*VideoGroupController 视频群组控制器 */
type VideoGroupController struct {
	UserController
}

/*VideoGroupResult 返回结果 */
type VideoGroupResult struct {
	Status       int                     `json:"status"`
	Msg          string                  `json:"msg"`
	Data         []models.SysVideoGroup       `json:"data" description:"视频群列表"`
}

/*AddVideoGroup 新增视频群组 */
// @Title 新增视频群组
// @Description 新增视频群组
// @Param data body models.SysVideoGroup true "请求参数"
// @Success 200 {object} controllers.c_group.VideoGroupResult
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /addVideoGroup [Post]
func (c *GroupController) AddVideoGroup() {
	var ob models.SysVideoGroup
	result := VideoGroupResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	res := models.AddVideoGroup(&ob)
	if res == 0 {
		result.Status = 0
		result.Msg = "会议室名称不能重复"
		c.Data["json"] = &result
		c.ServeJSON()
	}
	result.Status = 1
	result.Msg = "会议室创建成功"
	c.Data["json"] = &result
	c.ServeJSON()
}

/*SearchVideoGroup 搜索视频群组*/
// @Title 搜索视频群组
// @Description 搜索视频群组
// @Param data body models.SysVideoGroup true "请求参数"
// @Success 200 {object} controllers.c_group.VideoGroupResult
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /searchVideoGroup [Post]
func (c *GroupController) SearchVideoGroup() {
	var ob models.SysVideoGroup
	result := VideoGroupResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	resData := models.QuerVideoGroup(&ob)
	result.Status = 1
	result.Msg = "查询群组成功"
	result.Data = resData
	c.Data["json"] = &result
	c.ServeJSON()
}

/*GetVideoGroupList 获取群视频列表*/
// @Title 获取群视频列表
// @Description 获取群视频列表
// @Param data body models.SysVideoGroup true "请求参数"
// @Success 200 {object} controllers.c_group.VideoGroupResult
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /getVideoGroupList [Post]
func (c *GroupController) GetVideoGroupList() {
	var ob models.SysVideoGroup
	result := VideoGroupResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	resData := models.GetVideoGroupList(&ob)
	result.Status = 1
	result.Msg = "获取群列表成功"
	result.Data = resData
	c.Data["json"] = &result
	c.ServeJSON()
}

/*DelVideoGroup 删除会议室*/
// @Title 删除会议室
// @Description 删除会议室
// @Param data body models.SysVideoGroup true "请求参数"
// @Success 200 {object} controllers.c_group.VideoGroupResult
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /delVideoGroup [Post]
func (c *GroupController) DelVideoGroup() {
	var ob models.SysVideoGroup
	result := VideoGroupResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	res := models.DelVideoGroup(&ob)
	if res == 0 {
		result.Status = 0
		result.Msg = "删除失败"
	}
	result.Status = 1
	result.Msg = "删除会议室成功"
	c.Data["json"] = &result
	c.ServeJSON()
}
