package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

/*SysFriends 好友关系表信息 */
type SysFriends struct {
	Id         int       `orm:"auto" json:"id" required:"false" description:"好友关系id"`
	SourceUid  int       `json:"sourceUid" required:"false" description:"源用户id"`
	TargetUid  int       `json:"targetUid" required:"false" description:"目标用户id"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"createTime" required:"false" description:"创建时间"`
	SourceName string    `json:"sourceName" required:"false" description:"原用户名"`
	SourceMark string    `json:"sourceMark" required:"false" description:"原用户备注名"`
	TargetName string    `json:"targetName" required:"false" description:"目标用户名"`
	TargetMark string    `json:"targetMark" required:"false" description:"目标用户备注名"`
	Del        int       `json:"del" required:"false" description:"删除标记"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysFriends))
	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
}

/*AddFrends 新增好友关系*/
func AddFrends(friends *SysFriends) int64 {
	fmt.Println(friends)
	formData := SysFriends{}
	formData.SourceUid = friends.SourceUid
	formData.TargetUid = friends.TargetUid
	formData.SourceName = friends.SourceName
	formData.TargetName = friends.TargetName
	formData.SourceMark = "点击修改备注名"
	formData.TargetMark = "点击修改备注名"
	o := orm.NewOrm()
	if created, id, err := o.ReadOrCreate(&formData, "SourceUid", "TargetUid"); err == nil {
		if created {
			return id
		}
	}
	return 0
}

/*QuerFriends 按名称查找好友*/
func QuerFriends(friends *SysFriends) []SysFriends {
	var friendsData []SysFriends
	o := orm.NewOrm()
	o.QueryTable("sys_friends").Filter("sourceName", friends.SourceName).Limit(10).Offset(0).All(&friendsData)
	return friendsData
}

/*PutFriends 修改好友信息*/
func PutFriends(queryData *SysFriends) int64 {
	o := orm.NewOrm()
	sourceMark := queryData.SourceMark
	targetMark := queryData.TargetMark
	if len(queryData.SourceMark) == 0 {
		sourceMark = "点击修改备注名"
	}
	if len(queryData.TargetMark) == 0 {
		targetMark = "点击修改备注名"
	}
	num, _ := o.QueryTable("sys_friends").Filter("id", queryData.Id).Update(orm.Params{
		"del":     queryData.Del,
		"sourceMark": sourceMark,
		"targetMark": targetMark,
	})
	return num
}

/*GetFriendsList 查询全部好友列表（分页）*/
func GetFriendsList(queryData *SysFriends) []SysFriends {
	var friendsData []SysFriends
	o := orm.NewOrm()
	cond := orm.NewCondition()
	cond1 := cond.And("source_uid", queryData.Id).Or("target_uid", queryData.Id)

	qs := o.QueryTable("sys_friends")
	qs.SetCond(cond1).All(&friendsData)
	return friendsData
}
