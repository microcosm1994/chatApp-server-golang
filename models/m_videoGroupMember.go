package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

/*SysVideoGroupMember 群组成员 */
type SysVideoGroupMember struct {
	Id         int       `orm:"auto" json:"id" required:"false" description:"id"`
	User       *SysUser  `orm:"rel(fk)" json:"user" required:"false" description:"视频会议成员"`
	VideoGroup      *SysVideoGroup `orm:"rel(fk)" json:"videoGroup" required:"false" description:"视频会议信息"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"createTime" required:"false" description:"创建时间"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysVideoGroupMember))
	orm.Debug = false // 是否开启调试模式 调试模式下会打印出sql语句
}

/*AddVideoGroupMember 新增群成员*/
func AddVideoGroupMember(group *SysVideoGroupMember) int64 {
	formData := SysVideoGroupMember{}
	o := orm.NewOrm()
	// 多对一关系插入
	fkUser := SysUser{
		Id: group.User.Id,
	}
	fkVideoGroup := SysVideoGroup{
		Id: group.VideoGroup.Id,
	}
	formData.User = &fkUser
	formData.VideoGroup = &fkVideoGroup
	if created, id, err := o.ReadOrCreate(&formData, "user_id", "video_group_id"); err == nil {
		if created {
			return id
		}
	}
	return 0
}

/*GetVideoGroupUser 获取群成员*/
func GetVideoGroupUser(group *SysVideoGroup) []SysVideoGroupMember {
	var groupList []SysVideoGroupMember
	o := orm.NewOrm()
	o.QueryTable("sys_video_group_member").Filter("video_group_id", group.Id).RelatedSel().All(&groupList)
	return groupList
}

/*DelVideoGroupUser 移除群成员*/
func DelVideoGroupUser(group *SysVideoGroupMember) int64 {
	o := orm.NewOrm()
	formData := SysVideoGroupMember{
		Id: group.Id,
	}
	if num, err := o.Delete(&formData); err == nil {
		return num
	}
	return 0
}
