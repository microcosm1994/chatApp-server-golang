package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

/*SysUser 用户表信息 */
type SysUser struct {
	Id         int       `orm:"auto" json:"id" required:"false" description:"用户id"`
	NickName   string    `json:"nickName" required:"false" description:"用户昵称"`
	Phone      string    `orm:"unique" json:"phone" required:"false" description:"手机号"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"createTime" required:"false" description:"创建时间"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)" json:"updateTime" required:"false" description:"更新时间"`
}

/*SearchUser 手机号搜索用户 */
type SearchUser struct {
	Phone string `json:"phone" required:"true" description:"用户手机号"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysUser))
	orm.Debug = false // 是否开启调试模式 调试模式下会打印出sql语句
}

/*AddUser 新增用户*/
func AddUser(user *SearchUser) int64 {
	formData := SysUser{}
	formData.Phone = user.Phone
	formData.NickName = user.Phone
	o := orm.NewOrm()
	if created, id, err := o.ReadOrCreate(&formData, "phone"); err == nil {
		if created {
			return id
		}
	}
	return 0
}

/*QuerUser 按手机号查询用户列表*/
func QuerUser(user *SearchUser) []SysUser {
	var userData []SysUser
	o := orm.NewOrm()
	o.QueryTable("sys_user").Filter("phone", user.Phone).Limit(10).Offset(0).All(&userData)
	return userData
}

/*GetUserList 查询全部用户列表（分页）*/
func GetUserList(user *SysUser) []SysUser {
	var userData []SysUser
	o := orm.NewOrm()
	o.QueryTable("sys_user").Filter("phone", user.Phone).All(&userData)
	return userData
}
