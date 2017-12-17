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
	log := h.Server.GetLogger()
	players := h.Server.GetModule(map[uuid.UUID]sysinfo.PlayerInfo{}).(map[uuid.UUID]sysinfo.PlayerInfo)

	packet := &gamedefine.EnterRegionResultPacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("EnterRegionResultHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	players[packet.PlayerInfo.PlayerKey] = sysinfo.PlayerInfo{
		PlayerInfoBase: packet.PlayerInfo,
	}

	peerID, _ := uuid.FromString(packet.PeerID)
	clientPeer := h.Server.GetPeer(peerID)
	packet.PeerID = ""
	h.Server.SendPacket(clientPeer, packet)
	return true
}

// NewEnterRegionResultHandler is a constructor of EnterRegionResultHandler
func NewEnterRegionResultHandler(server ginterface.IGameServer) *EnterRegionResultHandler {
	ret := &EnterRegionResultHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, server)
	return ret
}
