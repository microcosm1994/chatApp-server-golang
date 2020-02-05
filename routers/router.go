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
		beego.NewNamespace("/v1",
			beego.NSNamespace("/user",
				beego.NSInclude(
					&controllers.UserController{},
				),
			),
			beego.NSNamespace("/friends",
				beego.NSInclude(
					&controllers.FriendsController{},
					&controllers.FriendsAskController{},
				),
			),
			beego.NSNamespace("/msg",
				beego.NSInclude(
					&controllers.MsgController{},
				),
			),
			beego.NSNamespace("/group",
				beego.NSInclude(
					&controllers.GroupController{},
					&controllers.GroupMsgController{},
				),
			),
		)
	beego.AddNamespace(ns)
}
