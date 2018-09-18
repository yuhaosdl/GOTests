package main

import (
	"GOTests/websocketserver/impl"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		}}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {

	// if conn, err := upgrader.Upgrade(w, r, nil); err == nil {
	// 	for {
	// 		if _, data, err := conn.ReadMessage(); err == nil {
	// 			if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
	// 				conn.Close()
	// 			}
	// 		} else {
	// 			conn.Close()
	// 		}
	// 	}
	// }
	if wsConn, err := upgrader.Upgrade(w, r, nil); err == nil {
		if conn, err := impl.InitConnection(wsConn); err == nil {
			for {
				if data, err := conn.ReadMessage(); err != nil {
					conn.Close()
				} else {
					if err := conn.WriteMessage(data); err != nil {
						conn.Close()
					}
				}
			}

		}
	}
}
func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}
