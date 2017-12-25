package ginterface

import (
	"github.com/satori/go.uuid"
	"golang.org/x/net/websocket"
)

// IGameServer is an interface of the game server
type IGameServer interface {
	INode

	// RegisterHandler registers the packet handler
	RegisterHandler(handler IGameHandler)
	// Start starts the server
	Start() bool
	// Stop stops the server
	Stop() bool

	// GetMasterPeer returns the game peer of the master server
	GetMasterPeer() IGamePeer
	// GetPeer returns the game peer of the peerID
	GetPeer(peerID uuid.UUID) IGamePeer

	// SendPacket sends packet to the connection
	SendPacket(peer IGamePeer, packet interface{}) bool
}

// IGameServerHook is an interface of hook of the game server
type IGameServerHook interface {
	// OnStart is called before Start()
	OnStart()
	// OnStopped is called at the end of Stop()
	OnStopped()
	// OnCreatePeer is called to create the custom peer
	OnCreatePeer(conn *websocket.Conn) IGamePeer
	// OnDefaultHandle is called when there is no corresponding packet handler
	OnDefaultHandle(peer IGamePeer, info string)
	// OnRegisterToMaster is called for registering to the master server
	OnRegisterToMaster()
}
