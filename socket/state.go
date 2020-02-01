package socket

import (
	"strconv"
	"github.com/googollee/go-socket.io"
)

/*State Socket方法与状态接口*/
type State interface{
	Add() bool
}

/*ConnectionPool Socket连接池*/
type ConnectionPool struct{
	Conn socketio.Conn
	User UserData
}

/*UserData Socket用户信息*/
type UserData struct{
	Id int
	Phone string
	NickName string
}

var pool map[string]ConnectionPool

/*Add 添加新的连接到连接池*/
func (c *ConnectionPool) Add(userData UserData, conn socketio.Conn) bool {
	var connData ConnectionPool
	connData.Conn = conn
	connData.User = userData
	pool[strconv.Itoa(userData.Id)] = connData
	return true
}



