package handler

import (
	"encoding/json"

	"../../base/ginterface"
	"../../base/handler"
	"../../gamecommon/gamedefine"
	uuid "github.com/satori/go.uuid"
)

// LoginResultHandler handles the result of player's login
type LoginResultHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *LoginResultHandler) Code() int {
	return gamedefine.LoginResult
}

// OnHandle is called when Handling the packet
func (h *LoginResultHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Node.GetLogger()

	packet := &gamedefine.LoginResultPacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("LoginResultHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	peerID, _ := uuid.FromString(packet.PeerID)
	clientPeer := h.Node.(ginterface.IGameServer).GetPeer(peerID)
	packet.PeerID = ""
	h.Node.(ginterface.IGameServer).SendPacket(clientPeer, packet)
	return true
}

// NewLoginResultHandler is a constructor of LoginResultHandler
func NewLoginResultHandler(node ginterface.INode) *LoginResultHandler {
	ret := &LoginResultHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, node)
	return ret
}
