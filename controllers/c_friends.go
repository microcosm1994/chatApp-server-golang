package controllers

import (
	"chatAppServer/models"
	"encoding/json"
	"fmt"
)

/*FriendsController 好友控制器 */
type FriendsController struct {
	MainController
}

/*SearchFriendsResult 返回结果 */
type SearchFriendsResult struct {
	Status int                 `json:"status"`
	Msg    string              `json:"msg"`
	Data   []models.SysFriends `json:"data" description:"好友数据"`
}

/*AddFriends 添加好友*/
// @Title 添加好友
// @Description 添加好友
// @Param data body models.SysFriends true "请求参数"
// @Success 200 {object} controllers.Result
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /addFriends [Post]
func (c *FriendsController) AddFriends() {
	var ob models.SysFriends
	result := SearchFriendsResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	resData := models.AddFrends(&ob)
	fmt.Println(resData)
	result.Status = 1
	result.Msg = "添加好友成功"
	c.Data["json"] = &result
	c.ServeJSON()
}

/*PutFriends 修改好友信息*/
// @Title 修改好友信息
// @Description 修改好友信息
// @Param data body models.SysFriends true "请求参数"
// @Success 200 {object} controllers.Result
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /putFriends [Post]
func (c *FriendsController) PutFriends() {
	var ob models.SysFriends
	result := SearchFriendsResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	resData := models.PutFriends(&ob)
	fmt.Println(resData)
	result.Status = 1
	result.Msg = "修改成功"
	c.Data["json"] = &result
	c.ServeJSON()
}

/*QueryFriends 按名称搜索好友*/
// @Title 按名称搜索好友
// @Description 按名称搜索好友
// @Param data body models.SearchFriends true "请求参数"
// @Success 200 {object} []models.SysFriends
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /queryFriend [Post]
func (c *FriendsController) QueryFriends() {
	var ob models.SysFriends
	result := SearchFriendsResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	resData := models.QuerFriends(&ob)
	fmt.Println(resData)
	result.Status = 1
	result.Msg = "查询好友成功"
	// 取出用户信息
	result.Data = resData
	c.Data["json"] = &result
	c.ServeJSON()
}

/*GetFriendsList 获取好友列表*/
// @Title 获取好友列表
// @Description 获取好友列表
// @Param data body {} true "请求参数"
// @Success 200 {object} []models.SysFriends
// @Failure 404 接口未找到
// @Failure 504 接口超时
// @router /getFriendsList [Post]
func (c *FriendsController) GetFriendsList() {
	var ob models.SysFriends
	result := SearchFriendsResult{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	resData := models.GetFriendsList(&ob)
	fmt.Println(resData)
	result.Status = 1
	result.Msg = "查询好友成功"
	// 取出用户信息
	result.Data = resData
	c.Data["json"] = &result
	c.ServeJSON()
}
