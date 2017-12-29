package handler

import (
	"encoding/json"

	"../../base/ginterface"
	"../../base/handler"
	"../../gamecommon/gamedefine"
	"../../servercommon/sysinfo"
	uuid "github.com/satori/go.uuid"
)

// EnterGameHandler handles the request of player's entering game
type EnterGameHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *EnterGameHandler) Code() int {
	return gamedefine.EnterGame
}

// OnHandle is called when Handling the packet
func (h *EnterGameHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Node.GetLogger()
	players := h.Node.(ginterface.IGameServer).GetModule(map[uuid.UUID]*sysinfo.PlayerInfo{}).(map[uuid.UUID]*sysinfo.PlayerInfo)
	games := h.Node.(ginterface.IGameServer).GetModule(map[int]ginterface.IGame{}).(map[int]ginterface.IGame)

	packet := &gamedefine.EnterGamePacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("EnterGameHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	_, ok := players[peer.GetPeerID()]
	if !ok {
		log.Error("EnterGameHandler.OnHandle(): no this player! peerID = %s\n", peer.GetPeerID().String())
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
