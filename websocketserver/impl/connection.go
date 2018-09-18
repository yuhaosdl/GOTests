package impl

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

//Connection ：封装websocketConnection
type Connection struct {
	wsConn    *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte
	isClosed  bool
	mutex     sync.Mutex
}

//InitConnection ：初始化Connection
func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn:    wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
	}
	//启动读协程
	go conn.readLoop()
	//启动写协程
	go conn.writeLoop()
	return
}

//api

//ReadMessage : 读信息
func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}

	return
}

//WriteMessage : 写信息
func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}

	return
}

//Close ： 关闭连接
func (conn *Connection) Close() {
	conn.wsConn.Close()
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

//内部实现

//readLoop : 循环读
func (conn *Connection) readLoop() {
	for {
		if _, data, err := conn.wsConn.ReadMessage(); err == nil {
			select {
			case conn.inChan <- data:
			case <-conn.closeChan:
				//closeChan关闭的时候
				conn.Close()
			}

		} else {
			conn.Close()
		}
	}

}

func (conn *Connection) writeLoop() {
	var (
		data []byte
	)
	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			conn.Close()
		}

		if err := conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			conn.Close()
		}
	}
}
