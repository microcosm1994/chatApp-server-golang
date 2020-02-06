package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

/*SysMsg 消息 */
type SysMsg struct {
	Id         int       `orm:"auto" json:"id" required:"false" description:"消息id"`
	FriendsId  int       `json:"friendsId" required:"false" description:"好友关系id"`
	TargetUid  int       `json:"targetUid" required:"false" description:"目标id"`
	SourceUid  int       `json:"sourceUid" required:"false" description:"源id"`
	Content    string    `json:"content" required:"false" description:"消息内容"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"createTime" required:"false" description:"创建时间"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysMsg))
	orm.Debug = false // 是否开启调试模式 调试模式下会打印出sql语句
}

/*AddMsg 新增消息*/
func AddMsg(msg *SysMsg) int64 {
	formData := SysMsg{}
	formData.FriendsId = msg.FriendsId
	formData.TargetUid = msg.TargetUid
	formData.SourceUid = msg.SourceUid
	formData.Content = msg.Content
	o := orm.NewOrm()
	id, err := o.Insert(&formData)
	if err == nil {
		return id
	}
	return 0
}


/*GetMsgList 获取用户消息*/
func GetMsgList(msg *SysMsg) []SysMsg {
	var msgList []SysMsg
	o := orm.NewOrm()
	o.QueryTable("sys_msg").Filter("friends_id", msg.FriendsId).All(&msgList)
	return msgList
}
