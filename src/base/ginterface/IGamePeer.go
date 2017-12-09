package ginterface

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/websocket"
)

// IGamePeer is an interface of the game peer
type IGamePeer interface {
	// GetPeerID gets the id of the peer
	GetPeerID() uuid.UUID
	// get the connection of the game peer
	GetConn() *websocket.Conn
	// OnConnected is called when connected
	OnConnected()
	// OnDisconnected is called when disconnected
	OnDisconnected()
}
