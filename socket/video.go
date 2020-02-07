package socket

import (
	"fmt"
	"strconv"

	socketio "github.com/googollee/go-socket.io"
)

/*Ops PeerOps*/
type Ops struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	FriendsId int    `json:"friendsId"`
	SourceId int    `json:"sourceId"`
	TargetId  int    `json:"targetId"`
	Data interface{} `json:"data"`
}

/*Result Peer返回结果*/
type Result struct {
	Type string      `json:"type"`
	Name string      `json:"name"`
	FriendsId int    `json:"friendsId"`
	SourceId int    `json:"sourceId"`
	TargetId  int    `json:"targetId"`
	Data interface{} `json:"data"`
}

/*VideoMount 挂载监听*/
func VideoMount(server socketio.Server, pool map[string]ConnectionPool) {
	JoinVideoRoom(server)
	SendVideoReq(server, pool)
	VideoReq(server)
	Ice(server)
	Offer(server)
	Answer(server)
}

/*JoinVideoRoom 加入视频房间*/
func JoinVideoRoom(server socketio.Server) {
	server.OnEvent("/socket.io/", "video_room", func(s socketio.Conn, msg Ops) error {
		// 生成房间
		room := "video_" + strconv.Itoa(msg.FriendsId)
		server.JoinRoom(room, s)
		return nil
	})
}

/*SendVideoReq 发送视频请求*/
func SendVideoReq(server socketio.Server, pool map[string]ConnectionPool) {
	server.OnEvent("/socket.io/", "video_send_req", func(s socketio.Conn, msg Ops) error {
		var result Result
		result.Name = msg.Name
		result.FriendsId = msg.FriendsId
		result.TargetId = msg.TargetId
		key := strconv.Itoa(msg.TargetId)
		pool[key].Conn.Emit("VideoReq", result)
		return nil
	})
}

/*VideoReq 视频请求*/
func VideoReq(server socketio.Server) {
	server.OnEvent("/socket.io/", "video_req", func(s socketio.Conn, msg Ops) error {
		var result Result
		result.Data = msg.Data
		fmt.Println(msg)
		fmt.Println(msg.Data)
		// 生成房间
		room := "video_" + strconv.Itoa(msg.FriendsId)
		server.BroadcastToRoom(room, "video_req_result", result)
		return nil
	})
}

/*Ice 转发ice*/
func Ice(server socketio.Server) {
	server.OnEvent("/socket.io/", "video_ice", func(s socketio.Conn, msg Ops, data interface{}) error {
		var result Result
		result.Type = msg.Type
		result.TargetId = msg.TargetId
		result.Data = data
		// 生成房间
		room := "video_" + strconv.Itoa(msg.FriendsId)
		server.BroadcastToRoom(room, "VideoIce", result)
		return nil
	})
}

/*Offer 转发信令*/
func Offer(server socketio.Server) {
	server.OnEvent("/socket.io/", "video_offer", func(s socketio.Conn, msg Ops, data interface{}) error {
		var result Result
		result.Type = msg.Type
		result.TargetId = msg.TargetId
		result.Data = data
		// 生成房间
		room := "video_" + strconv.Itoa(msg.FriendsId)
		server.BroadcastToRoom(room, "VideoOffer", result)
		return nil
	})
}

/*Answer 转发应答*/
func Answer(server socketio.Server) {
	server.OnEvent("/socket.io/", "video_answer", func(s socketio.Conn, msg Ops, data interface{}) error {
		var result Result
		result.Type = msg.Type
		result.TargetId = msg.TargetId
		result.Data = data
		// 生成房间
		room := "video_" + strconv.Itoa(msg.FriendsId)
		server.BroadcastToRoom(room, "VideoAnswer", result)
		return nil
	})
}
