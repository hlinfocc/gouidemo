package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // 允许来自任何源的连接，实际项目中应根据情况调整
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	for {
		// 读取消息
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading from WebSocket:", err)
			break
		}
		log.Printf("Received: %s", message)

		// 回复消息
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error writing to WebSocket:", err)
			break
		}
	}
}

func StartWebsocket() {
	http.HandleFunc("/wsclient", wsEndpoint)
	fmt.Println("WebSocket server is running on :32388")
	log.Fatal(http.ListenAndServe(":32388", nil))
}
