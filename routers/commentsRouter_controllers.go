package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["chatAppServer/controllers:FriendsAskController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:FriendsAskController"],
        beego.ControllerComments{
            Method: "AddFriendsAsk",
            Router: `/addFriendsAsk`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:FriendsAskController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:FriendsAskController"],
        beego.ControllerComments{
            Method: "GetFriendsAskList",
            Router: `/getFriendsAskList`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:FriendsAskController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:FriendsAskController"],
        beego.ControllerComments{
            Method: "PutFriendsAsk",
            Router: `/putFriendsAsk`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:FriendsController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:FriendsController"],
        beego.ControllerComments{
            Method: "AddFriends",
            Router: `/addFriends`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:FriendsController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:FriendsController"],
        beego.ControllerComments{
            Method: "GetFriendsList",
            Router: `/getFriendsList`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:FriendsController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:FriendsController"],
        beego.ControllerComments{
            Method: "PutFriends",
            Router: `/putFriends`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:FriendsController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:FriendsController"],
        beego.ControllerComments{
            Method: "QueryFriends",
            Router: `/queryFriend`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:MsgController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:MsgController"],
        beego.ControllerComments{
            Method: "AddMsg",
            Router: `/addMsg`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:MsgController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:MsgController"],
        beego.ControllerComments{
            Method: "GetMsgList",
            Router: `/getMsgList`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

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

    beego.GlobalControllerRouter["chatAppServer/controllers:UserController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:UserController"],
        beego.ControllerComments{
            Method: "SearchUser",
            Router: `/searchUser`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
