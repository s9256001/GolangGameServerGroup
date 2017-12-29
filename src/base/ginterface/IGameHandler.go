package ginterface

// IGameHandler is an interface of the packet handler
type IGameHandler interface {
	// Code is the associated packet command
	Code() int
	// Handle handles the packet
	Handle(peer IGamePeer, info string) bool
}

// IGameHandlerHook is an interface of hook of the packet handler
type IGameHandlerHook interface {
	// OnHandle is called when Handling the packet
	OnHandle(peer IGamePeer, info string) bool
}
