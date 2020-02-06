package socket

import (
	"strconv"
	"chatAppServer/models"
	socketio "github.com/googollee/go-socket.io"
)

/*GroupData Socket组信息*/
type GroupData struct {
	Id       int
	Phone    string
	NickName string
}

/*GroupMsgMount 挂载监听*/
func GroupMsgMount(server socketio.Server) {
	JoinGroupRoom(server)
	SendGroupMsg(server)
}

/*JoinGroupRoom 加入群组房间*/
func JoinGroupRoom(server socketio.Server) {
	server.OnEvent("/socket.io/", "JoinGroupRoom", func(s socketio.Conn, msg models.SysGroupMsg) error {
		groupRoom := "g_" + strconv.Itoa(msg.GroupId)
		server.JoinRoom(groupRoom, s)
		return nil
	})
}

/*SendGroupMsg 发送群组消息*/
func SendGroupMsg(server socketio.Server) {
	server.OnEvent("/socket.io/", "SendGroupMessage", func(s socketio.Conn, msg models.SysGroupMsg) error {
		// 生成房间
		groupRoom := "g_" + strconv.Itoa(msg.GroupId)
		server.BroadcastToRoom(groupRoom, "GroupRoomMsg", msg)
		models.AddGroupMsg(&msg)
		return nil
	})
}
