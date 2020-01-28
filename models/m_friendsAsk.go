package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

/*SysFriends 好友关系表信息 */
type SysFriendsAsk struct {
	Id         int       `orm:"auto" json:"id" required:"false" description:"好友关系id"`
	SourceUid  int       `json:"sourceUid" required:"false" description:"源用户id"`
	TargetUid  int       `json:"targetUid" required:"false" description:"目标用户id"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"createTime" required:"false" description:"创建时间"`
	SourceName string    `json:"sourceName" required:"false" description:"原用户名"`
	SourceMark string    `json:"sourceMark" required:"false" description:"原用户备注名"`
	TargetName  string    `json:"targetName" required:"false" description:"目标用户名"`
	TargetMark  string    `json:"targetMark" required:"false" description:"目标用户备注名"`
	Del  int    `orm:"default(0)" json:"del" required:"false" description:"软删除标记0：未删除，1：已删除"`
	IsAgree  int    `orm:"default(0)" json:"isAgree" required:"false" description:"是否同意请求0：未同意，1：同意，2：拒绝	"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysFriendsAsk))
	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
}

/*AddFrendsAsk 新增好友关系*/
func AddFrendsAsk(friends *SysFriendsAsk) int64 {
	fmt.Println(friends)
	formData := SysFriendsAsk{}
	formData.SourceUid = friends.SourceUid
	formData.TargetUid = friends.TargetUid
	formData.SourceName = friends.SourceName
	formData.SourceMark = friends.SourceMark
	formData.TargetName = friends.TargetName
	formData.TargetMark = friends.TargetMark
	o := orm.NewOrm()
	if created, id, err := o.ReadOrCreate(&formData, "SourceUid", "TargetUid"); err == nil {
		if created {
			return id
		}
	}
	return 0
}

/*GetFriendsAskList 查询全部好友请求列表*/
func GetFriendsAskList(queryData *SysFriendsAsk) []SysFriendsAsk {
	var friendsAskData []SysFriendsAsk
	o := orm.NewOrm()
	o.QueryTable("sys_friends_ask").Filter("del", queryData.Del).All(&friendsAskData)
	return friendsAskData
}
