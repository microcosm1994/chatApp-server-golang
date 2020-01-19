package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["chatAppServer/controllers:UserController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/login`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
