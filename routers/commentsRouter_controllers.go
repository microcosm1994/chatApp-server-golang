package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["chatAppServer/controllers:UserController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:UserController"],
        beego.ControllerComments{
            Method: "CheckCode",
            Router: `/checkCode`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:UserController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
