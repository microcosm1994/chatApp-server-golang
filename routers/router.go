package routers

import (
	"chatAppServer/controllers"
	"github.com/astaxie/beego"
)


func init() {
	ns :=
		beego.NewNamespace("/user",
			beego.NSRouter("/getMessageCode", &controllers.UserController{}, "*:GetMessageCode"),
			beego.NSRouter("/checkingCode", &controllers.UserController{}, "*:CheckingCode"),
		)
	beego.AddNamespace(ns)
}
