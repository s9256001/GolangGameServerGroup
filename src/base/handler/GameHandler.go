package handler

import (
	"../ginterface"
)

// GameHandler is an abstract class of the packet handler
// It fits the interface of IGameHandler, and depends on the derived type to implement the interface of IGameHandlerHook
type GameHandler struct {
	ginterface.IGameHandlerHook // hook

	Node ginterface.INode // node
}

// Code is the associated packet command
func (h *GameHandler) Code() int {
	return 0
}

// Handle handles the packet
func (h *GameHandler) Handle(peer ginterface.IGamePeer, info string) bool {
	return h.OnHandle(peer, info)
}

// NewGameHandler is a constructor of GameHandler
func NewGameHandler(hook ginterface.IGameHandlerHook, node ginterface.INode) *GameHandler {
	ret := &GameHandler{
		IGameHandlerHook: hook,

		Node: node,
	}
	return ret
}
