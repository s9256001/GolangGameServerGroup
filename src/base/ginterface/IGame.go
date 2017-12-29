package ginterface

// IGame is an interface of a game
type IGame interface {
	INode // base interface

	// GameID gets the game ID of the game
	GameID() int
	// RegisterHandler registers the packet handler
	RegisterHandler(handler IGameHandler)
	// Init initlizes the game
	Init(setting interface{}) bool
	// Release releases the game
	Release() bool
	// HandlePacket handles the packet
	HandlePacket(peer IGamePeer, info string)
	// todo
	// GetServer returns the game server
	GetServer() IGameServer
}

// IGameHook is an interface of hook of the game server
type IGameHook interface {
	// OnInit is called when Init()
	OnInit(setting interface{}) bool
	// OnRelease is called when Release()
	OnRelease() bool
	// OnDefaultHandle is called when there is no corresponding packet handler
	OnDefaultHandle(peer IGamePeer, info string)
}
