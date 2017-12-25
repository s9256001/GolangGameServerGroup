package handler

import (
	"encoding/json"

	"../../base/ginterface"
	"../../base/handler"
	"../../gamecommon/gamedefine"
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
	games := h.Node.(ginterface.IGameServer).GetModule(map[int]ginterface.IGame{}).(map[int]ginterface.IGame)

	packet := &gamedefine.EnterGamePacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("EnterGameHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	game, ok := games[packet.GameID]
	if !ok {
		log.Error("EnterGameHandler.OnHandle(): no this game! gameID = %d\n", packet.GameID)
		return false
	}

	game.HandlePacket(peer, info)
	return true
}

// NewEnterGameHandler is a constructor of EnterGameHandler
func NewEnterGameHandler(node ginterface.INode) *EnterGameHandler {
	ret := &EnterGameHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, node)
	return ret
}
