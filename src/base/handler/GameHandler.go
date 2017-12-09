package handler

import (
	"../ginterface"
)

// GameHandler is an abstract class of the packet handler
// It fits the interface of IGameHandler, and depends on the derived type to implement the interface of IGameHandlerHook
type GameHandler struct {
	ginterface.IGameHandlerHook // hook

	Server ginterface.IGameServer // server
}

// Handle handles the packet
func (h *GameHandler) Handle(peer ginterface.IGamePeer, info string) bool {
	return h.OnHandle(peer, info)
}

// NewGameHandler is a constructor of GameHandler
func NewGameHandler(hook ginterface.IGameHandlerHook, server ginterface.IGameServer) *GameHandler {
	ret := &GameHandler{
		IGameHandlerHook: hook,

		Server: server,
	}
	return ret
}
