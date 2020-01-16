package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type SysUser struct {
	Id         int    `json:"id"`
	NickName   string `json:"nickName"`
	Phone      string `json:"phone"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SysUser))
	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
}

func AddUser(user *SysUser) string {
	fmt.Println(user)
	println(user.Phone)
	return user.Phone
}

func QuerUser(user *SysUser) bool {
	var userData []SysUser
	// 获取 QueryBuilder 对象. 需要指定数据库驱动参数。
	// 第二个返回值是错误对象，在这里略过
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("*").From("sys_user").
		Where("phone = ?").
		OrderBy("update_time").Desc().
		Limit(10).Offset(0)
	// 导出 SQL 语句
	sql := qb.String()
	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(sql, user.Phone).QueryRows(&userData)

	// o := orm.NewOrm()
	// o.QueryTable("sysUser").Filter("phone", user.Phone).All(&userData)
	fmt.Println(userData)
	return true
}
