// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

import (
	"chatAppServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns :=
		beego.NewNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		)
	beego.AddNamespace(ns)
}
