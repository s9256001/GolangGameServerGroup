package handler

import (
	"encoding/json"

	"../../base/ginterface"
	"../../base/handler"
	"../../gamecommon/gamedefine"
)

// LoginHandler handles player's login
type LoginHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *LoginHandler) Code() int {
	return gamedefine.Login
}

// OnHandle is called when Handle()
func (h *LoginHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Node.GetLogger()

	packet := &gamedefine.LoginPacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("LoginHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	packet.PeerID = peer.GetPeerID().String()
	h.Node.(ginterface.IGameServer).SendPacket(h.Node.(ginterface.IGameServer).GetMasterPeer(), packet)
	return true
}

// NewLoginHandler is a constructor of LoginHandler
func NewLoginHandler(node ginterface.INode) *LoginHandler {
	ret := &LoginHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, node)
	return ret
}
