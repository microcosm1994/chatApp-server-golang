package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

/*SysFriends 好友请求表信息 */
type SysFriendsAsk struct {
	Id         int       `orm:"auto" json:"id" required:"false" description:"好友关系id"`
	SourceUid  int       `json:"sourceUid" required:"false" description:"源用户id"`
	TargetUid  int       `json:"targetUid" required:"false" description:"目标用户id"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"createTime" required:"false" description:"创建时间"`
	SourceName string    `json:"sourceName" required:"false" description:"原用户名"`
	TargetName string    `json:"targetName" required:"false" description:"目标用户名"`
	Del        int       `orm:"default(0)" json:"del" required:"false" description:"软删除标记0：未删除，1：已删除"`
	IsAgree    int       `orm:"default(0)" json:"isAgree" required:"false" description:"是否同意请求0：未同意，1：同意，2：拒绝	"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysFriendsAsk))
	orm.Debug = false // 是否开启调试模式 调试模式下会打印出sql语句
}

/*AddFrendsAsk 新增好友请求*/
func AddFrendsAsk(friends *SysFriendsAsk) int64 {
	var data []SysFriendsAsk
	formData := SysFriendsAsk{}
	formData.SourceUid = friends.SourceUid
	formData.TargetUid = friends.TargetUid
	formData.SourceName = friends.SourceName
	formData.TargetName = friends.TargetName
	o := orm.NewOrm()
	// 交叉验证
	cond := orm.NewCondition()
	cond1 := cond.And("source_uid", friends.SourceUid).And("target_uid", friends.TargetUid)
	cond2 := cond.And("source_uid", friends.TargetUid).And("target_uid", friends.SourceUid)
	condResult := cond.AndCond(cond1).OrCond(cond2)
	qs := o.QueryTable("sys_friends_ask")
	qs.SetCond(condResult).All(&data)
	if len(data) == 0 {
		id, err := o.Insert(&formData)
		if err == nil {
			return id
		}
	}
	return 0
}

/*PutFriendsAsk 修改好友请求*/
func PutFriendsAsk(queryData *SysFriendsAsk) int64 {
	o := orm.NewOrm()
	num, _ := o.QueryTable("sys_friends_ask").Filter("id", queryData.Id).Update(orm.Params{
		"del":     queryData.Del,
		"isAgree": queryData.IsAgree,
	})
	return num
}

/*GetFriendsAskList 查询全部好友请求列表*/
func GetFriendsAskList(queryData *SysFriendsAsk) []SysFriendsAsk {
	var friendsAskData []SysFriendsAsk
	o := orm.NewOrm()
	o.QueryTable("sys_friends_ask").Filter("target_uid", queryData.TargetUid).Filter("del", queryData.Del).All(&friendsAskData)
	return friendsAskData
}
