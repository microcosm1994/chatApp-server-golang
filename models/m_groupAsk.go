package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

/*SysGroupAsk 群组 */
type SysGroupAsk struct {
	Id         int       `orm:"auto" json:"id" required:"false" description:"群id"`
	GroupMpId  int       `json:"groupMpId" required:"false" description:"群主id"`
	IsAgree    int       `orm:"default(0)" json:"isAgree" required:"false" description:"是否同意0：无状态，1：同意，2：拒绝"`
	Del        int       `orm:"default(0)" json:"del" required:"false" description:"软删除0：不删除，1：删除"`
	User       *SysUser  `orm:"rel(fk)" json:"user" required:"false" description:"申请人信息"`
	Group      *SysGroup `orm:"rel(fk)" json:"group" required:"false" description:"组信息"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"createTime" required:"false" description:"创建时间"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysGroupAsk))
	orm.Debug = false // 是否开启调试模式 调试模式下会打印出sql语句
}

/*AddGroupAsk 新增群组申请*/
func AddGroupAsk(group *SysGroupAsk) int64 {
	formData := SysGroupAsk{}
	formData.GroupMpId = group.GroupMpId
	// 多对一插入
	fkUser := SysUser{
		Id: group.User.Id,
	}
	fkGroup := SysGroup{
		Id: group.Group.Id,
	}
	formData.Group = &fkGroup
	formData.User = &fkUser
	o := orm.NewOrm()
	if created, id, err := o.ReadOrCreate(&formData, "user_id", "group_id"); err == nil {
		if created {
			return id
		}
	}
	return 0
}

/*PutGroupAsk 修改好友请求*/
func PutGroupAsk(group *SysGroupAsk) int64 {
	o := orm.NewOrm()
	num, _ := o.QueryTable("sys_group_ask").Filter("id", group.Id).Update(orm.Params{
		"del":     group.Del,
		"isAgree": group.IsAgree,
	})
	return num
}

/*GetGroupAsk 获取群申请列表*/
func GetGroupAsk(group *SysGroupAsk) []SysGroupAsk {
	var groupList []SysGroupAsk
	o := orm.NewOrm()
	o.QueryTable("sys_group_ask").Filter("group_mp_id", group.GroupMpId).RelatedSel().All(&groupList)
	return groupList
}

/*DelGroupAsk 删除群申请*/
func DelGroupAsk(group *SysGroupAsk) int64 {
	o := orm.NewOrm()
	var ask SysGroupAsk
	// 获取ask id
	o.QueryTable("sys_group_ask").Filter("group_id", group.Group.Id).Filter("user_id", group.User.Id).All(&ask)
	formData := SysGroupAsk{
		Id: ask.Id,
	}
	if num, err := o.Delete(&formData); err == nil {
		return num
	}
	return 0
}
