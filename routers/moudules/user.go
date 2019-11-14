package routers

import (
	"chatAppServer/controllers"
	"context"

	"github.com/astaxie/beego"
)

func init() {
	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/shop",
				beego.NSGet("/:id", func(ctx *context.Context) {
					ctx.Output.Body([]byte("shopinfo"))
				}),
			),
			beego.NSNamespace("/order",
				beego.NSGet("/:id", func(ctx *context.Context) {
					ctx.Output.Body([]byte("orderinfo"))
				}),
			),
			beego.NSNamespace("/crm",
				beego.NSGet("/:id", func(ctx *context.Context) {
					ctx.Output.Body([]byte("crminfo"))
				}),
			),
		)
	// beego.Router("/user/getMessageCode", &controllers.UserController{})
}
