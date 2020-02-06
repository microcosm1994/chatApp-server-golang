package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

/*SysGroupMsg 消息 */
type SysGroupMsg struct {
	Id         int       `orm:"auto" json:"id" required:"false" description:"消息id"`
	GroupId    int       `json:"groupId" required:"false" description:"组id"`
	UserId     int       `json:"userId" required:"false" description:"用户id"`
	Content    string    `json:"content" required:"false" description:"消息内容"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"createTime" required:"false" description:"创建时间"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysGroupMsg))
	orm.Debug = false // 是否开启调试模式 调试模式下会打印出sql语句
}

/*AddGroupMsg 新增组消息*/
func AddGroupMsg(msg *SysGroupMsg) int64 {
	formData := SysGroupMsg{}
	formData.GroupId = msg.GroupId
	formData.UserId = msg.UserId
	formData.Content = msg.Content
	o := orm.NewOrm()
	id, err := o.Insert(&formData)
	if err == nil {
		return id
	}
	return 0
}

/*GetGroupMsgList 获取组消息*/
func GetGroupMsgList(msg *SysGroupMsg) []SysGroupMsg {
	var groupMsgList []SysGroupMsg
	o := orm.NewOrm()
	o.QueryTable("sys_group_msg").Filter("group_id", msg.GroupId).All(&groupMsgList)
	return groupMsgList
}
