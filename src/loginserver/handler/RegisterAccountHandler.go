package handler

import (
	"encoding/json"

	"../../base/ginterface"
	"../../base/handler"
	"../../gamecommon/gamedefine"
	"../../servercommon/sysdefine"
)

// RegisterAccountHandler handles the request of the registration of the player's account
type RegisterAccountHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *RegisterAccountHandler) Code() int {
	return gamedefine.RegisterAccount
}

// OnHandle is called when Handling the packet
func (h *RegisterAccountHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Node.GetLogger()

	response := gamedefine.NewRegisterAccountResultPacket()
	response.Result = sysdefine.Failed

	packet := &gamedefine.RegisterAccountPacket{}

	defer func() {
		if response.Result == sysdefine.OK {
			h.Node.(ginterface.IGameServer).SendPacket(h.Node.(ginterface.IGameServer).GetMasterPeer(), packet)
		} else {
			h.Node.(ginterface.IGameServer).SendPacket(peer, response)
		}
	}()

	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("RegisterAccountHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	if packet.Account == "" || len(packet.Account) > gamedefine.MaxAccountLength || len(packet.Account) < gamedefine.MinAccountLength {
		response.Result = gamedefine.RegisterAccountInvalidAccount
		return false
	}
	if packet.Password == "" || len(packet.Password) > gamedefine.MaxPasswordLength || len(packet.Password) < gamedefine.MinPasswordLength {
		response.Result = gamedefine.RegisterAccountInvalidPassword
		return false
	}

	packet.PeerID = peer.GetPeerID().String()
	response.Result = sysdefine.OK
	return true
}

// NewRegisterAccountHandler is a constructor of RegisterAccountHandler
func NewRegisterAccountHandler(node ginterface.INode) *RegisterAccountHandler {
	ret := &RegisterAccountHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, node)
	return ret
}
