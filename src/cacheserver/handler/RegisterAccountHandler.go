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
	log := h.Node.(ginterface.IGameServer).GetLogger()

	response := gamedefine.NewRegisterAccountResultPacket()
	response.Result = sysdefine.Failed

	defer h.Node.(ginterface.IGameServer).SendPacket(peer, response)

	packet := &gamedefine.RegisterAccountPacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("RegisterAccountHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	response.PeerID = packet.PeerID
	response.Result = sysdefine.OK
	return true
}

// NewRegisterAccountHandler is a constructor of RegisterAccountHandler
func NewRegisterAccountHandler(node ginterface.INode) *RegisterAccountHandler {
	ret := &RegisterAccountHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, node)
	return ret
}
