package handler

import (
	"encoding/json"

	"../../base/ginterface"
	"../../base/handler"
	"../../gamecommon/gamedefine"
	uuid "github.com/satori/go.uuid"
)

// LoginResultHandler handles the registration of the player's account
type LoginResultHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *LoginResultHandler) Code() int {
	return gamedefine.LoginResult
}

// OnHandle is called when Handle()
func (h *LoginResultHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Server.GetLogger()

	packet := &gamedefine.LoginResultPacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("LoginResultHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	peerID, _ := uuid.FromString(packet.PeerID)
	clientPeer := h.Server.GetPeer(peerID)
	packet.PeerID = ""
	h.Server.SendPacket(clientPeer, packet)
	return true
}

// NewLoginResultHandler is a constructor of LoginResultHandler
func NewLoginResultHandler(server ginterface.IGameServer) *LoginResultHandler {
	ret := &LoginResultHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, server)
	return ret
}
