package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

/*SysGroupMember 群组成员 */
type SysGroupMember struct {
	Id         int       `orm:"auto" json:"id" required:"false" description:"id"`
	UserId     int       `json:"userId" required:"false" description:"成员id(用户id)"`
	Group      *SysGroup `orm:"rel(fk)" json:"group" required:"false" description:"组信息"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"createTime" required:"false" description:"创建时间"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysGroupMember))
	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
}

/*AddGroupMember 新增群成员*/
func AddGroupMember(group *SysGroupMember) int64 {
	formData := SysGroupMember{}
	formData.UserId = group.UserId
	o := orm.NewOrm()
	// 多对一关系插入
	fkGroup := SysGroup{
		Id: group.Group.Id,
	}
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
	o.QueryTable("sys_group_member").Filter("user_id", group.UserId).RelatedSel().All(&groupList)
	fmt.Println(groupList)
	return groupList
}
