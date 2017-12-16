package handler

import (
	"encoding/json"

	"../../base/ginterface"
	"../../base/handler"
	"../../gamecommon/gamedefine"
	uuid "github.com/satori/go.uuid"
)

// RegisterAccountResultHandler handles the registration of the player's account
type RegisterAccountResultHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *RegisterAccountResultHandler) Code() int {
	return gamedefine.RegisterAccountResult
}

// OnHandle is called when Handle()
func (h *RegisterAccountResultHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Server.GetLogger()

	packet := &gamedefine.RegisterAccountResultPacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("RegisterAccountResultHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}
	log.Debug("RegisterAccountResultHandler.OnHandle(): code = %d, result = %v\n", h.Code(), packet.Result)

	peerID, _ := uuid.FromString(packet.PeerID)
	clientPeer := h.Server.GetPeer(peerID)
	packet.PeerID = ""
	h.Server.SendPacket(clientPeer, packet)
	return true
}

// NewRegisterAccountResultHandler is a constructor of RegisterAccountResultHandler
func NewRegisterAccountResultHandler(server ginterface.IGameServer) *RegisterAccountResultHandler {
	ret := &RegisterAccountResultHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, server)
	return ret
}
