package ginterface

import (
	"golang.org/x/net/websocket"
)

// IGameServer is an interface of the game server
type IGameServer interface {
	// RegisterHandler registers the packet handler
	RegisterHandler(handler IGameHandler)
	// Start starts the server
	Start() bool
	// Stop stops the server
	Stop() bool

	// GetLogger returns the game logger
	GetLogger() IGameLogger
	// GetMasterPeer returns the game peer of the master server
	GetMasterPeer() IGamePeer
	// GetModule returns the specific module to resolve import cycle
	GetModule(module interface{}) interface{}

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
