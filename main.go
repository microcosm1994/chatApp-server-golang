package main

import (
	_ "chatAppServer/routers"
	"chatAppServer/filters"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"chatAppServer/utils"
)

func init() {
	//注册数据库
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:113655@tcp(localhost:3306)/chatapp_golang?charset=utf8")
	// 过滤器，token验证
	beego.InsertFilter("/*", beego.BeforeExec, filters.TokenFilter)
	// 建立redis连接
	utils.InitRedis()
}

func main() {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	// swagger自动化生成文档
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
