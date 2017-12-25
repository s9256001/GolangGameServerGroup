package handler

import (
	"encoding/json"

	"../../base/ginterface"
	"../../base/handler"
	"../../gamecommon/gamedefine"
	"../../servercommon/sysdefine"
)

// RegisterAccountHandler handles the registration of the player's account
type RegisterAccountHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *RegisterAccountHandler) Code() int {
	return gamedefine.RegisterAccount
}

// OnHandle is called when Handle()
func (h *RegisterAccountHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Node.GetLogger()

	response := gamedefine.NewRegisterAccountResultPacket()
	response.Result = sysdefine.OK

	packet := &gamedefine.RegisterAccountPacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("RegisterAccountHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	if packet.Account == "" {
		response.Result = gamedefine.RegisterAccountInvalidAccount
		h.Node.(ginterface.IGameServer).SendPacket(peer, response)
		return false
	}
	if packet.Password == "" {
		response.Result = gamedefine.RegisterAccountInvalidPassword
		h.Node.(ginterface.IGameServer).SendPacket(peer, response)
		return false
	}

	packet.PeerID = peer.GetPeerID().String()
	h.Node.(ginterface.IGameServer).SendPacket(h.Node.(ginterface.IGameServer).GetMasterPeer(), packet)
	return true
}

// NewRegisterAccountHandler is a constructor of RegisterAccountHandler
func NewRegisterAccountHandler(node ginterface.INode) *RegisterAccountHandler {
	ret := &RegisterAccountHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, node)
	return ret
}
