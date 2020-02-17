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

    beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"],
        beego.ControllerComments{
            Method: "AddGroup",
            Router: `/addGroup`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"],
        beego.ControllerComments{
            Method: "AddGroupAsk",
            Router: `/addGroupAsk`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"],
        beego.ControllerComments{
            Method: "AddGroupMember",
            Router: `/addGroupMember`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"],
        beego.ControllerComments{
            Method: "AddVideoGroup",
            Router: `/addVideoGroup`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"],
        beego.ControllerComments{
            Method: "AddVideoGroupMember",
            Router: `/addVideoGroupMember`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"],
        beego.ControllerComments{
            Method: "DelGroupUser",
            Router: `/delGroupUser`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"],
        beego.ControllerComments{
            Method: "DelVideoGroup",
            Router: `/delVideoGroup`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"],
        beego.ControllerComments{
            Method: "DelVideoGroupUser",
            Router: `/delVideoGroupUser`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"],
        beego.ControllerComments{
            Method: "GetGroupAskList",
            Router: `/getGroupAskList`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"],
        beego.ControllerComments{
            Method: "GetGroupList",
            Router: `/getGroupList`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"],
        beego.ControllerComments{
            Method: "GetGroupUserList",
            Router: `/getGroupUserList`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"],
        beego.ControllerComments{
            Method: "GetVideoGroupList",
            Router: `/getVideoGroupList`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"],
        beego.ControllerComments{
            Method: "GetVideoGroupUserList",
            Router: `/getVideoGroupUserList`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"],
        beego.ControllerComments{
            Method: "PutGroupAsk",
            Router: `/putGroupAsk`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"],
        beego.ControllerComments{
            Method: "SearchGroup",
            Router: `/searchGroup`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:GroupController"],
        beego.ControllerComments{
            Method: "SearchVideoGroup",
            Router: `/searchVideoGroup`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:GroupMsgController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:GroupMsgController"],
        beego.ControllerComments{
            Method: "AddGroupMsg",
            Router: `/`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatAppServer/controllers:GroupMsgController"] = append(beego.GlobalControllerRouter["chatAppServer/controllers:GroupMsgController"],
        beego.ControllerComments{
            Method: "GetGroupMsgList",
            Router: `/getGroupMsgList`,
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
