package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

/*SysVideoGroup 群组 */
type SysVideoGroup struct {
	Id         int       `json:"id" required:"false" description:"群id"`
	GroupMpId  int       `json:"groupMpId" required:"false" description:"群主id"`
	GroupName  string    `orm:"unique" json:"groupName" required:"false" description:"群名称"`
	GroupDes   string    `json:"groupDes" required:"false" description:"群说明"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"createTime" required:"false" description:"创建时间"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysVideoGroup))
	orm.Debug = false // 是否开启调试模式 调试模式下会打印出sql语句
}

/*AddVideoGroup 新增视频群组*/
func AddVideoGroup(group *SysVideoGroup) int64 {
	formData := SysVideoGroup{}
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

/*GetVideoGroupList 获取视频群列表*/
func GetVideoGroupList(group *SysVideoGroup) []SysVideoGroup {
	var groupList []SysVideoGroup
	o := orm.NewOrm()
	o.QueryTable("sys_video_group").All(&groupList)
	return groupList
}

/*QuerVideoGroup 按名称查询视频群列表*/
func QuerVideoGroup(group *SysVideoGroup) []SysVideoGroup {
	var groupList []SysVideoGroup
	o := orm.NewOrm()
	o.QueryTable("sys_video_group").Filter("group_name", group.GroupName).Limit(10).Offset(0).All(&groupList)
	return groupList
}

/*DelVideoGroup 删除会议室*/
func DelVideoGroup(group *SysVideoGroup) int64 {
	o := orm.NewOrm()
	formData := SysVideoGroup{
		Id: group.Id,
	}
	if num, err := o.Delete(&formData); err == nil {
		return num
	}
	return 0
}
