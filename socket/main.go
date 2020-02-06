package socket

import (
	"log"
	"net/http"
	"strconv"
	socketio "github.com/googollee/go-socket.io"
)

/*Socket socket接口*/
type Socket interface {}

/*ConnectionPool Socket连接池*/
type ConnectionPool struct {
	Conn socketio.Conn
	User UserData
}

/*UserData Socket组信息*/
type UserData struct {
	Id       int
	Phone    string
	NickName string
}

/*Server socket实例*/
var Server socketio.Server
var pool = make(map[string]ConnectionPool)

/*NewServer 建立socket*/
func NewServer() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		return nil
	})
	server.OnEvent("/socket.io/", "ConnStatus", func(s socketio.Conn, msg UserData) error {
		Server = *server
		key := strconv.Itoa(msg.Id)
		pool[key] = ConnectionPool{
			Conn: s,
			User: msg,
		}
		pool[key].Conn.Emit("ConnStatus", 1)
		return nil
	})
	server.OnError("/socket.io/", func(s socketio.Conn, e error) {
	})
	MsgMount(*server)
	GroupMsgMount(*server)
	VideoMount(*server, pool)
	go server.Serve()
	defer server.Close()
	http.Handle("/", server)
	log.Fatal(http.ListenAndServe(":9090", nil))
	log.Println("Serving at localhost:9090...")
}
