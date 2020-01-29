package controllers

import (
	"chatAppServer/models"
	"encoding/json"
	"fmt"
)

/*FriendsAskController 好友请求控制器 */
type FriendsAskController struct {
	MainController
}

/*FriendsAskResult 返回结果 */
type FriendsAskResult struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data []models.SysFriendsAsk `json:"data"`
}

/*AddFriendsAsk 添加新的好友请求*/
// @Title 添加新的好友请求
// @Description 添加新的好友请求
// @Param data body models.SysFriendsAsk true "请求参数"
// @Success 200 {object} controllers.FriendsAskResult
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /addFriendsAsk [Post]
func (c *FriendsAskController) AddFriendsAsk() {
	var ob models.SysFriendsAsk
	result := FriendsAskResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	resId := models.AddFrendsAsk(&ob)
	fmt.Println(resId)
	if resId == 0 {
		result.Msg = "您已发送过好友申请,请等待对方同意。"
	} else {
		result.Msg = "已发送好友请求"
	}
	result.Status = 1
	c.Data["json"] = &result
	c.ServeJSON()
}

/*PutFriendsAsk 修改好友请求*/
// @Title 修改好友请求
// @Description 修改好友请求
// @Param data body models.SysFriendsAsk true "请求参数"
// @Success 200 {object} controllers.Result
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /putFriendsAsk [Post]
func (c *FriendsAskController) PutFriendsAsk() {
	var ob models.SysFriendsAsk
	result := FriendsAskResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	num := models.PutFriendsAsk(&ob)
	fmt.Println("num",num)
	if num == 0 {
		result.Msg = "修改失败"
	} else {
		result.Msg = "修改成功"
	}
	result.Status = 1
	c.Data["json"] = &result
	c.ServeJSON()
}

/*GetFriendsAskList 获取好友请求列表*/
// @Title 获取好友列表
// @Description 获取好友列表
// @Param data body models.m_friendsAsk.SysFriendsAsk true "请求参数"
// @Success 200 {object} []models.SysFriendsAsk
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /getFriendsAskList [Post]
func (c *FriendsAskController) GetFriendsAskList() {
	var ob models.SysFriendsAsk
	result := FriendsAskResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	resData := models.GetFriendsAskList(&ob)
	fmt.Println(resData)
	result.Status = 1
	result.Msg = "获取好友请求列表成功"
	// 取出用户信息
	result.Data = resData
	c.Data["json"] = &result
	c.ServeJSON()
}
