package socket

import (
	"fmt"
	"strconv"
	"chatAppServer/models"
	socketio "github.com/googollee/go-socket.io"
)

/*UserData Socket用户信息*/
type UserData struct {
	Id       int
	Phone    string
	NickName string
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
		fmt.Println(msg)
		// 生成房间
		room := strconv.Itoa(msg.FriendsId)
		server.BroadcastToRoom(room, "RoomMsg", msg)
		models.AddMsg(&msg)
		return nil
	})
}
