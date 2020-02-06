package socket

import (
	"strconv"
	"chatAppServer/models"
	socketio "github.com/googollee/go-socket.io"
)

/*MsgMount 挂载监听*/
func MsgMount(server socketio.Server) {
	JoinRoom(server)
	SendMsg(server)
}

/*JoinRoom 加入房间*/
func JoinRoom(server socketio.Server) {
	server.OnEvent("/socket.io/", "JoinRoom", func(s socketio.Conn, msg models.SysMsg) error {
		server.JoinRoom(strconv.Itoa(msg.FriendsId), s)
		return nil
	})
}

/*SendMsg 发送消息*/
func SendMsg(server socketio.Server) {
	server.OnEvent("/socket.io/", "SendMessage", func(s socketio.Conn, msg models.SysMsg) error {
		// 生成房间
		room := strconv.Itoa(msg.FriendsId)
		server.BroadcastToRoom(room, "RoomMsg", msg)
		models.AddMsg(&msg)
		return nil
	})
}
