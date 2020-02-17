package socket

import (
	"fmt"
	"strconv"

	socketio "github.com/googollee/go-socket.io"
)

/*VideoGroupOps PeerOps*/
type VideoGroupOps struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	VideoGroupId int    `json:"videoGroupId"`
	UserId int    `json:"userId"`
	Data interface{} `json:"data"`
}

/*VideoGroupMount 挂载监听*/
func VideoGroupMount(server socketio.Server) {
	JoinVideoGroupRoom(server)
	LeaveVideoGroupRoom(server)
	GroupIce(server)
	GroupOffer(server)
	GroupAnswer(server)
}

/*JoinVideoGroupRoom 加入视频房间*/
func JoinVideoGroupRoom(server socketio.Server) {
	server.OnEvent("/socket.io/", "video_group_room", func(s socketio.Conn, msg VideoGroupOps) error {
		var result VideoGroupOps
		result.UserId = msg.UserId
		result.VideoGroupId = msg.VideoGroupId
		// 生成房间名称
		room := "video_group_" + strconv.Itoa(msg.VideoGroupId)
		server.JoinRoom(room, s)
		server.BroadcastToRoom(room, "VideoGroupNotice", result)
		fmt.Println(server.Rooms())
		fmt.Println(s.Rooms())
		return nil
	})
}

/*LeaveVideoGroupRoom 离开视频房间*/
func LeaveVideoGroupRoom(server socketio.Server) {
	server.OnEvent("/socket.io/", "video_group_room_leave", func(s socketio.Conn, msg VideoGroupOps) error {
		// 生成房间名称
		room := "video_group_" + strconv.Itoa(msg.VideoGroupId)
		server.LeaveRoom(room, s)
		return nil
	})
}

/*GroupIce 转发ice*/
func GroupIce(server socketio.Server) {
	server.OnEvent("/socket.io/", "video_group_ice", func(s socketio.Conn, msg VideoGroupOps, data interface{}) error {
		var result VideoGroupOps
		result.Type = msg.Type
		result.UserId = msg.UserId
		result.Data = data
		// 生成房间
		room := "video_group_" + strconv.Itoa(msg.VideoGroupId)
		server.BroadcastToRoom(room, "VideoGroupIce", result)
		return nil
	})
}

/*GroupOffer 转发信令*/
func GroupOffer(server socketio.Server) {
	server.OnEvent("/socket.io/", "video_group_offer", func(s socketio.Conn, msg VideoGroupOps, data interface{}) error {
		var result VideoGroupOps
		result.UserId = msg.UserId
		result.Data = data
		// 生成房间
		room := "video_group_" + strconv.Itoa(msg.VideoGroupId)
		server.BroadcastToRoom(room, "VideoGroupOffer", result)
		return nil
	})
}

/*GroupAnswer 转发应答*/
func GroupAnswer(server socketio.Server) {
	server.OnEvent("/socket.io/", "video_group_answer", func(s socketio.Conn, msg VideoGroupOps, data interface{}) error {
		var result VideoGroupOps
		result.UserId = msg.UserId
		result.Data = data
		// 生成房间
		room := "video_group_" + strconv.Itoa(msg.VideoGroupId)
		server.BroadcastToRoom(room, "VideoGroupAnswer", result)
		return nil
	})
}
