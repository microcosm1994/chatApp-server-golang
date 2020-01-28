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
	TargetName  string    `json:"targetName" required:"false" description:"目标用户名"`
	TargetMark  string    `json:"targetMark" required:"false" description:"目标用户备注名"`
}

/*SearchFriends 名称搜索用户 */
type SearchFriends struct {
	SourceName string `json:"sourceName" required:"true" description:"原用户名"`
	SourceMark string    `json:"sourceMark" required:"false" description:"原用户备注名"`
	TargetName string `json:"targetName" required:"true" description:"目标用户名"`
	TargetMark  string    `json:"targetMark" required:"false" description:"目标用户备注名"`
	Offset int `json:"offset"  required:"true" description:"偏移量"`
	Limit int `json:"limit"  required:"true" description:"数据条数"`
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
	formData.SourceMark = friends.SourceMark
	formData.TargetName = friends.TargetName
	formData.TargetMark = friends.TargetMark
	o := orm.NewOrm()
	id, err := o.Insert(&formData)
	if err == nil {
		return id
	}
	return 0
}

/*QuerFriends 按名称查找好友*/
func QuerFriends(friends *SearchFriends) []SysFriends {
	var friendsData []SysFriends
	o := orm.NewOrm()
	o.QueryTable("sys_friends").Filter("sourceName", friends.SourceName).Limit(10).Offset(0).All(&friendsData)
	return friendsData
}

/*GetFriendsList 查询全部好友列表（分页）*/
func GetFriendsList(queryData *SearchFriends) []SysFriends {
	var friendsData []SysFriends
	o := orm.NewOrm()
	o.QueryTable("sys_friends").All(&friendsData)
	return friendsData
}
