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

// Code is the associated packet command
func (h *GameHandler) Code() int {
	return 0
}

// Handle handles the packet
func (h *GameHandler) Handle(peer ginterface.IGamePeer, info string) bool {
	log := h.Server.GetLogger()

	log.Trace("GameHandler.Handle: code = %d, info = %s\n", h.Code(), info)
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
