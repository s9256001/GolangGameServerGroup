package handler

import (
	"encoding/json"

	"../../base/ginterface"
	"../../base/handler"
	"../../gamecommon/gamedefine"
	"../../servercommon/sysdefine"
)

// EnterRegionHandler handles the request of player's entering region
type EnterRegionHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *EnterRegionHandler) Code() int {
	return gamedefine.EnterRegion
}

// OnHandle is called when Handling the packet
func (h *EnterRegionHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Node.GetLogger()

	response := gamedefine.NewEnterRegionResultPacket()
	response.Result = sysdefine.Failed

	packet := &gamedefine.EnterRegionPacket{}

	defer func() {
		if response.Result == sysdefine.OK {
			h.Node.(ginterface.IGameServer).SendPacket(h.Node.(ginterface.IGameServer).GetMasterPeer(), packet)
		} else {
			h.Node.(ginterface.IGameServer).SendPacket(peer, response)
		}
	}()

	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("EnterRegionHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	packet.PeerID = peer.GetPeerID().String()
	response.Result = sysdefine.OK
	return true
}

// NewEnterRegionHandler is a constructor of EnterRegionHandler
func NewEnterRegionHandler(node ginterface.INode) *EnterRegionHandler {
	ret := &EnterRegionHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, node)
	return ret
}
