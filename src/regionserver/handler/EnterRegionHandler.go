package handler

import (
	"encoding/json"

	"../../base/ginterface"
	"../../base/handler"
	"../../gamecommon/gamedefine"
)

// EnterRegionHandler handles player's entering region
type EnterRegionHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *EnterRegionHandler) Code() int {
	return gamedefine.EnterRegion
}

// OnHandle is called when Handle()
func (h *EnterRegionHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Server.GetLogger()

	packet := &gamedefine.EnterRegionPacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("EnterRegionHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	packet.PeerID = peer.GetPeerID().String()
	h.Server.SendPacket(h.Server.GetMasterPeer(), packet)
	return true
}

// NewEnterRegionHandler is a constructor of EnterRegionHandler
func NewEnterRegionHandler(server ginterface.IGameServer) *EnterRegionHandler {
	ret := &EnterRegionHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, server)
	return ret
}
