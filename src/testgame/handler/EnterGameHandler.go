package handler

import (
	"encoding/json"

	"../../base/ginterface"
	"../../base/handler"
	"../../gamecommon/gamedefine"
	"../../servercommon/sysdefine"
)

// EnterGameHandler handles player's entering game
type EnterGameHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *EnterGameHandler) Code() int {
	return gamedefine.EnterGame
}

// OnHandle is called when Handle()
func (h *EnterGameHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Node.GetLogger()

	packet := &gamedefine.EnterGamePacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("EnterGameHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}
	result := gamedefine.NewEnterGameResultPacket()
	result.GameID = packet.GameID
	result.Result = sysdefine.OK

	h.Node.(ginterface.IGame).GetServer().SendPacket(peer, result)
	return true
}

// NewEnterGameHandler is a constructor of EnterGameHandler
func NewEnterGameHandler(node ginterface.INode) *EnterGameHandler {
	ret := &EnterGameHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, node)
	return ret
}
