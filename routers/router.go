package routers

import (
	"chatAppServer/controllers"
	"github.com/astaxie/beego"
)


func init() {
	ns :=
		beego.NewNamespace("/user",
			beego.NSRouter("/login", &controllers.UserController{}, "*:Login"),
		)
	beego.AddNamespace(ns)
}
