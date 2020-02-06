package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

/*SysGroupMember 群组成员 */
type SysGroupMember struct {
	Id         int       `orm:"auto" json:"id" required:"false" description:"id"`
	User       *SysUser  `orm:"rel(fk)" json:"user" required:"false" description:"群组成员"`
	Group      *SysGroup `orm:"rel(fk)" json:"group" required:"false" description:"组信息"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"createTime" required:"false" description:"创建时间"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysGroupMember))
	orm.Debug = false // 是否开启调试模式 调试模式下会打印出sql语句
}

/*AddGroupMember 新增群成员*/
func AddGroupMember(group *SysGroupMember) int64 {
	formData := SysGroupMember{}
	o := orm.NewOrm()
	// 多对一关系插入
	fkUser := SysUser{
		Id: group.User.Id,
	}
	fkGroup := SysGroup{
		Id: group.Group.Id,
	}
	formData.User = &fkUser
	formData.Group = &fkGroup
	if created, id, err := o.ReadOrCreate(&formData, "user_id", "group_id"); err == nil {
		if created {
			return id
		}
	}
	return 0
}

/*GetGroupMember 获取群列表*/
func GetGroupMember(group *SysGroupMember) []SysGroupMember {
	var groupList []SysGroupMember
	o := orm.NewOrm()
	o.QueryTable("sys_group_member").Filter("user_id", group.User.Id).RelatedSel().All(&groupList)
	return groupList
}

/*GetGroupUser 获取群成员*/
func GetGroupUser(group *SysGroup) []SysGroupMember {
	var groupList []SysGroupMember
	o := orm.NewOrm()
	o.QueryTable("sys_group_member").Filter("group_id", group.Id).RelatedSel().All(&groupList)
	return groupList
}

/*DelGroupUser 移除群成员*/
func DelGroupUser(group *SysGroupMember) int64 {
	o := orm.NewOrm()
	formData := SysGroupMember{
		Id: group.Id,
	}
	if num, err := o.Delete(&formData); err == nil {
		return num
	}
	return 0
}
