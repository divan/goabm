package ui

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WSServer struct {
	upgrader websocket.Upgrader
	hub      []*websocket.Conn
}

func NewWSServer() *WSServer {
	ws := &WSServer{
		upgrader: websocket.Upgrader{},
	}
	return ws
}

type WSResponse struct {
	Type MsgType         `json:"type"`
	Data json.RawMessage `json:"data,omitempty"`
}

type WSRequest struct {
	Cmd WSCommand `json:"cmd"`
}

type MsgType string
type WSCommand string

// WebSocket response types
const (
	RespData MsgType = "data"
)

// WebSocket commands
const (
	CmdUpdate WSCommand = "update"
)

func (ws *WSServer) Handle(w http.ResponseWriter, r *http.Request) {
	c, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()

	ws.hub = append(ws.hub, c)

	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			break
		}
	}
}

func (ws *WSServer) sendMsg(c *websocket.Conn, msg *WSResponse) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Println("marshall:", err)
		return
	}

	err = c.WriteMessage(1, data)
	if err != nil {
		log.Println("write:", err)
		c = nil
		return
	}
}

func (ws *WSServer) broadcastData(data []byte) {
	msg := &WSResponse{
		Type: RespData,
		Data: json.RawMessage(data),
	}
	for i := 0; i < len(ws.hub); i++ {
		if ws.hub[i] == nil {
			continue
		}
		ws.sendMsg(ws.hub[i], msg)
	}
}
