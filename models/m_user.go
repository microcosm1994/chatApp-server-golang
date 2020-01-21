package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

/*SysUser 用户表信息 */
type SysUser struct {
	Id         int       `orm:"auto" json:"id" required:"false" description:"用户id"`
	NickName   string    `orm:"default(aa)" json:"nickName" required:"false" description:"用户昵称"`
	Phone      string    `orm:"unique" json:"phone" required:"false" description:"手机号"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"createTime" required:"false" description:"创建时间"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)" json:"updateTime" required:"false" description:"更新时间"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysUser))
	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
}

/*AddUser 新增用户*/
func AddUser(user *SysUser) int64 {
	formData := SysUser{}
	formData.Phone = user.Phone
	o := orm.NewOrm()
	id, err := o.Insert(&formData)
	if err == nil {
		return id
	}
	return 0
}

/*QuerUser 按手机号查询用户列表*/
func QuerUser(user *SysUser) []SysUser {
	var userData []SysUser
	o := orm.NewOrm()
	o.QueryTable("sys_user").Filter("phone", user.Phone).All(&userData)
	return userData
}
