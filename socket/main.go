package socket

import (
	"fmt"
	"log"
	"net/http"
	"github.com/googollee/go-socket.io"
)

func init() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.Emit("res", "ssssssss")
		fmt.Println("Url:", s.URL())
		fmt.Println("connected:", s.ID())
		return nil
	})
	go server.Serve()
	defer server.Close()
	http.Handle("/", server)
	log.Fatal(http.ListenAndServe(":9090", nil))
	log.Println("Serving at localhost:9090...")
}

