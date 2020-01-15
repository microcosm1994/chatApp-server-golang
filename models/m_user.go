package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	ID int `json:"id"`
	NickName string `json:"nickName"`
	Phone string `json:"phone"`
	CreateTime string `json:"createTime"`
	UptadeTime string `json:"updateTime"`
}

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(new(User))
}

func AddUser(user *User) string {
	println(user)
	println(user.Phone)
	return user.Phone
}

func QuerUser(user *User) string {
	println(user.NickName)
	println(user.ID)
	println(user.Phone)
	return user.NickName
}
