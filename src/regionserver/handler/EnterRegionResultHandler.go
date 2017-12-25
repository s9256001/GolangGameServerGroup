package handler

import (
	"encoding/json"

	"../../base/ginterface"
	"../../base/handler"
	"../../gamecommon/gamedefine"
	"../../servercommon/sysinfo"
	uuid "github.com/satori/go.uuid"
)

// EnterRegionResultHandler handles player's entering region
type EnterRegionResultHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *EnterRegionResultHandler) Code() int {
	return gamedefine.EnterRegionResult
}

// OnHandle is called when Handle()
func (h *EnterRegionResultHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Node.GetLogger()
	players := h.Node.(ginterface.IGameServer).GetModule(map[uuid.UUID]sysinfo.PlayerInfo{}).(map[uuid.UUID]sysinfo.PlayerInfo)

	packet := &gamedefine.EnterRegionResultPacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("EnterRegionResultHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	players[packet.PlayerInfo.PlayerKey] = sysinfo.PlayerInfo{
		PlayerInfoBase: packet.PlayerInfo,
	}

	peerID, _ := uuid.FromString(packet.PeerID)
	clientPeer := h.Node.(ginterface.IGameServer).GetPeer(peerID)
	packet.PeerID = ""
	h.Node.(ginterface.IGameServer).SendPacket(clientPeer, packet)
	return true
}

// NewEnterRegionResultHandler is a constructor of EnterRegionResultHandler
func NewEnterRegionResultHandler(node ginterface.INode) *EnterRegionResultHandler {
	ret := &EnterRegionResultHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, node)
	return ret
}
