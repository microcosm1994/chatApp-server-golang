package controllers

import (
	"chatAppServer/models"
	"encoding/json"
	"strconv"
)

/*GroupController 群组控制器 */
type GroupController struct {
	UserController
}

/*GroupResult 返回结果 */
type GroupResult struct {
	Status       int                     `json:"status"`
	Msg          string                  `json:"msg"`
	Data         []models.SysGroup       `json:"data" description:"群列表"`
	GroupList    []models.SysGroupMember `json:"groupList" description:"群列表"`
	GroupAskList []models.SysGroupAsk    `json:"groupAskList" description:"群申请列表"`
}

/*AddGroup 新增群组 */
// @Title 新增群组
// @Description 新增群组
// @Param data body models.SysGroup true "请求参数"
// @Success 200 {object} controllers.c_group.GroupResult
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /addGroup [Post]
func (c *GroupController) AddGroup() {
	var ob models.SysGroup
	result := GroupResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	res := models.AddGroup(&ob)
	if res == 0 {
		result.Status = 0
		result.Msg = "群名称不能重复"
		c.Data["json"] = &result
		c.ServeJSON()
	} else {
		var member models.SysGroupMember
		// int64转int
		strRes := strconv.FormatInt(res, 10)
		intRes, _ := strconv.Atoi(strRes)
		group := models.SysGroup{
			Id: intRes,
		}
		user := models.SysUser{
			Id: ob.GroupMpId,
		}
		member.User = &user
		member.Group = &group
		// 新增群成员
		models.AddGroupMember(&member)
	}
	result.Status = 1
	result.Msg = "群组创建成功"
	c.Data["json"] = &result
	c.ServeJSON()
}

/*AddGroupMember 新增群成员 */
// @Title 新增群成员
// @Description 新增群成员
// @Param data body models.SysGroupMember true "请求参数"
// @Success 200 {object} controllers.c_group.GroupResult
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /addGroupMember [Post]
func (c *GroupController) AddGroupMember() {
	var ob models.SysGroupMember
	group := models.SysGroup{}
	user := models.SysUser{}
	ob.User = &user
	ob.Group = &group
	result := GroupResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	res := models.AddGroupMember(&ob)
	if res == 0 {
		result.Status = 0
		result.Msg = "用户已经加入群组"
		c.Data["json"] = &result
		c.ServeJSON()
	}
	result.Status = 1
	result.Msg = "添加群成员成功"
	c.Data["json"] = &result
	c.ServeJSON()
}

/*AddGroupAsk 新增群申请 */
// @Title 新增群申请
// @Description 新增群申请
// @Param data body models.SysGroupAsk true "请求参数"
// @Success 200 {object} controllers.c_group.GroupResult
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /addGroupAsk [Post]
func (c *GroupController) AddGroupAsk() {
	var ob models.SysGroupAsk
	group := models.SysGroup{}
	user := models.SysUser{}
	ob.Group = &group
	ob.User = &user
	result := GroupResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	res := models.AddGroupAsk(&ob)
	if res == 0 {
		result.Status = 0
		result.Msg = "您已发送过申请，请等待对方同意"
		c.Data["json"] = &result
		c.ServeJSON()
	}
	result.Status = 1
	result.Msg = "发送申请成功，请等待对方同意"
	c.Data["json"] = &result
	c.ServeJSON()
}

/*SearchGroup 搜索群组*/
// @Title 搜索群组
// @Description 搜索群组
// @Param data body models.SysGroup true "请求参数"
// @Success 200 {object} controllers.c_group.SysGroup
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /searchGroup [Post]
func (c *GroupController) SearchGroup() {
	var ob models.SysGroup
	result := GroupResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	resData := models.QuerGroup(&ob)
	result.Status = 1
	result.Msg = "查询群组成功"
	result.Data = resData
	c.Data["json"] = &result
	c.ServeJSON()
}

/*GetGroupList 获取群组列表*/
// @Title 获取群组列表
// @Description 获取群组列表
// @Param data body models.SysGroupMember true "请求参数"
// @Success 200 {object} controllers.c_group.GroupResult
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /getGroupList [Post]
func (c *GroupController) GetGroupList() {
	var ob models.SysGroupMember
	user := models.SysUser{}
	ob.User = &user
	result := GroupResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	resData := models.GetGroupMember(&ob)
	result.Status = 1
	result.Msg = "获取群列表成功"
	result.GroupList = resData
	c.Data["json"] = &result
	c.ServeJSON()
}

/*GetGroupUserList 获取群成员*/
// @Title 获取群成员
// @Description 获取群成员
// @Param data body models.SysGroupMember true "请求参数"
// @Success 200 {object} controllers.c_group.GroupResult
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /getGroupUserList [Post]
func (c *GroupController) GetGroupUserList() {
	var ob models.SysGroup
	result := GroupResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	resData := models.GetGroupUser(&ob)
	result.Status = 1
	result.Msg = "获取群成员成功"
	result.GroupList = resData
	c.Data["json"] = &result
	c.ServeJSON()
}

/*GetGroupAskList 获取群申请列表*/
// @Title 获取群申请列表
// @Description 获取群申请列表
// @Param data body models.SysGroupAsk true "请求参数"
// @Success 200 {object} controllers.c_group.GroupResult
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /getGroupAskList [Post]
func (c *GroupController) GetGroupAskList() {
	var ob models.SysGroupAsk
	result := GroupResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	resData := models.GetGroupAsk(&ob)
	result.Status = 1
	result.Msg = "获取群申请列表成功"
	result.GroupAskList = resData
	c.Data["json"] = &result
	c.ServeJSON()
}

/*PutGroupAsk 修改群申请*/
// @Title 修改群申请
// @Description 修改群申请
// @Param data body models.SysGroupAsk true "请求参数"
// @Success 200 {object} controllers.c_group.GroupResult
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /putGroupAsk [Post]
func (c *GroupController) PutGroupAsk() {
	var ob models.SysGroupAsk
	result := GroupResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	res := models.PutGroupAsk(&ob)
	if res == 0 {
		result.Status = 0
		result.Msg = "修改群申请失败"
	}
	result.Status = 1
	result.Msg = "修改群申请成功"
	c.Data["json"] = &result
	c.ServeJSON()
}

/*DelGroupUser 移除群成员*/
// @Title 移除群成员
// @Description 移除群成员
// @Param data body models.SysGroupMember true "请求参数"
// @Success 200 {object} controllers.c_group.GroupResult
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /delGroupUser [Post]
func (c *GroupController) DelGroupUser() {
	var ob models.SysGroupMember
	user := models.SysUser{}
	group := models.SysGroup{}
	ob.User = &user
	ob.Group = &group
	result := GroupResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	res := models.DelGroupUser(&ob)
	if res == 0 {
		result.Status = 0
		result.Msg = "移除群成员失败"
	} else {
		ask := models.SysGroupAsk{}
		askUser := models.SysUser{
			Id: ob.User.Id,
		}
		askGroup := models.SysGroup{
			Id: ob.Group.Id,
		}
		ask.User = &askUser
		ask.Group = &askGroup
		models.DelGroupAsk(&ask)
	}
	result.Status = 1
	result.Msg = "移除群成员成功"
	c.Data["json"] = &result
	c.ServeJSON()
}
