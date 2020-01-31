package socket

import (
	"fmt"
	"log"
	"net/http"
	"github.com/googollee/go-socket.io"
)

func NewServer() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.Emit("res", "ssssssss")
		fmt.Println("Url:", s.URL())
		fmt.Println("connected:", s.ID())
		s.SetContext("1111")
		return nil
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})
	go server.Serve()
	defer server.Close()
	http.Handle("/socket.io/", server)
	http.ListenAndServe(":8000", nil)
	log.Println("Serving at localhost:8000...")
}

