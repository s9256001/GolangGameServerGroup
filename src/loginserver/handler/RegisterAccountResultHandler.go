package handler

import (
	"encoding/json"

	"../../base/ginterface"
	"../../base/handler"
	"../../gamecommon/gamedefine"
	uuid "github.com/satori/go.uuid"
)

// RegisterAccountResultHandler handles the result of the registration of the player's account
type RegisterAccountResultHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *RegisterAccountResultHandler) Code() int {
	return gamedefine.RegisterAccountResult
}

// OnHandle is called when Handling the packet
func (h *RegisterAccountResultHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Node.GetLogger()

	packet := &gamedefine.RegisterAccountResultPacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("RegisterAccountResultHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	peerID, _ := uuid.FromString(packet.PeerID)
	clientPeer := h.Node.(ginterface.IGameServer).GetPeer(peerID)
	packet.PeerID = ""
	h.Node.(ginterface.IGameServer).SendPacket(clientPeer, packet)
	return true
}

// NewRegisterAccountResultHandler is a constructor of RegisterAccountResultHandler
func NewRegisterAccountResultHandler(node ginterface.INode) *RegisterAccountResultHandler {
	ret := &RegisterAccountResultHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, node)
	return ret
}
