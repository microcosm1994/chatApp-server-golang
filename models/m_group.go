package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

/*SysGroup 群组 */
type SysGroup struct {
	Id         int       `json:"id" required:"false" description:"群id"`
	GroupMpId  int       `json:"groupMpId" required:"false" description:"群主id"`
	GroupName  string    `orm:"unique" json:"groupName" required:"false" description:"群名称"`
	GroupDes   string    `json:"groupDes" required:"false" description:"群说明"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"createTime" required:"false" description:"创建时间"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysGroup))
	orm.Debug = false // 是否开启调试模式 调试模式下会打印出sql语句
}

/*AddGroup 新增群组*/
func AddGroup(group *SysGroup) int64 {
	formData := SysGroup{}
	formData.GroupMpId = group.GroupMpId
	formData.GroupName = group.GroupName
	formData.GroupDes = group.GroupDes
	o := orm.NewOrm()
	if created, id, err := o.ReadOrCreate(&formData, "group_name"); err == nil {
		if created {
			return id
		}
	}
	return 0
}

/*QuerGroup 按群名称查询群列表*/
func QuerGroup(group *SysGroup) []SysGroup {
	var groupList []SysGroup
	o := orm.NewOrm()
	o.QueryTable("sys_group").Filter("group_name", group.GroupName).Limit(10).Offset(0).All(&groupList)
	return groupList
}
