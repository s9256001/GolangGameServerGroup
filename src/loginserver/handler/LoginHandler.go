package handler

import (
	"encoding/json"

	"../../base/ginterface"
	"../../base/handler"
	"../../gamecommon/gamedefine"
)

// LoginHandler handles the registration of the player's account
type LoginHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *LoginHandler) Code() int {
	return gamedefine.Login
}

// OnHandle is called when Handle()
func (h *LoginHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Server.GetLogger()

	packet := &gamedefine.LoginPacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("LoginHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	packet.PeerID = peer.GetPeerID().String()
	h.Server.SendPacket(h.Server.GetMasterPeer(), packet)
	return true
}

// NewLoginHandler is a constructor of LoginHandler
func NewLoginHandler(server ginterface.IGameServer) *LoginHandler {
	ret := &LoginHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, server)
	return ret
}
